package di

import (
	"reflect"
	"testing"
)

// --- Mocks and Types for Testing ---

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Calculator struct {
	Config *Config
}

type Config struct {
	Factor int
}

// --- Factory Functions ---

func NewConfig() *Config {
	return &Config{Factor: 10}
}

func NewCircle() *Circle {
	return &Circle{Radius: 5}
}

func NewCalculator(cfg *Config) *Calculator {
	return &Calculator{Config: cfg}
}

// --- Helper: Internal Reset ---
// Since providerRegistry is private, we implement a Reset function
// to ensure test isolation.
func resetRegistry() {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	providerRegistry = make(map[reflect.Type][]*Provider)
}

// --- Test Cases ---

func TestDI_RegisterAndResolve(t *testing.T) {
	resetRegistry()

	Register(NewConfig)

	cfg := Resolve[*Config]()
	if cfg == nil {
		t.Fatal("Expected *Config, got nil")
	}
	if cfg.Factor != 10 {
		t.Errorf("Expected factor 10, got %d", cfg.Factor)
	}
}

func TestDI_Singleton(t *testing.T) {
	resetRegistry()

	Singleton(NewConfig)

	inst1 := Resolve[*Config]()
	inst2 := Resolve[*Config]()

	// Compara endereços de memória para garantir que é a mesma instância
	if inst1 != inst2 {
		t.Error("Singleton failed: instances are different, memory addresses do not match")
	}
}

func TestDI_Transient(t *testing.T) {
	resetRegistry()

	Register(NewConfig)

	inst1 := Resolve[*Config]()
	inst2 := Resolve[*Config]()

	// Devem ser instâncias diferentes (endereços diferentes)
	if inst1 == inst2 {
		t.Error("Transient failed: instances should have different memory addresses")
	}
}

func TestDI_RegisterAs(t *testing.T) {
	resetRegistry()

	// Registers *Circle bound to Shape interface
	RegisterAs[Shape](NewCircle)

	shape := Resolve[Shape]()
	if shape == nil {
		t.Fatal("Failed to resolve Shape interface")
	}

	expectedArea := 78.5
	if shape.Area() != expectedArea {
		t.Errorf("Expected area %f, got %f", expectedArea, shape.Area())
	}
}

func TestDI_NestedDependencies(t *testing.T) {
	resetRegistry()

	// Register dependencies
	Singleton(NewConfig)
	Register(NewCalculator) // NewCalculator depends on *Config

	calc := Resolve[*Calculator]()

	if calc.Config == nil {
		t.Fatal("Dependency *Config was not automatically injected into Calculator")
	}

	if calc.Config.Factor != 10 {
		t.Errorf("Injected *Config has incorrect value: %d", calc.Config.Factor)
	}
}

func TestDI_ResolveAll(t *testing.T) {
	resetRegistry()

	// Register multiple shapes
	RegisterAs[Shape](func() *Circle { return &Circle{Radius: 1} })
	RegisterAs[Shape](func() *Circle { return &Circle{Radius: 2} })

	shapes := ResolveAll[Shape]()

	if len(shapes) != 2 {
		t.Errorf("Expected 2 registered shapes, got %d", len(shapes))
	}
}

func TestDI_Panics(t *testing.T) {
	t.Run("Panic on non-function factory", func(t *testing.T) {
		resetRegistry()
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should have panicked when registering a string instead of a function")
			}
		}()
		Register("not a function")
	})

	t.Run("Panic on multi-return factory", func(t *testing.T) {
		resetRegistry()
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should have panicked when factory returns more than one value")
			}
		}()
		multiReturnFactory := func() (*Config, error) { return &Config{}, nil }
		Register(multiReturnFactory)
	})
}
