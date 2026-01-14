package mut

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Mutable is the generic interface for use in your code.
type Mutable[T any] interface {
	Get() T
	Set(T)
	Dirty() bool
}

// Mut is the structure that tracks the value state.
type Mut[T any] struct {
	dirty bool
	value T
}

// Get returns the stored value.
func (m *Mut[T]) Get() T { return m.value }

// Dirty returns true if the value has been modified.
func (m *Mut[T]) Dirty() bool { return m.dirty }

// Set updates the value and marks it as dirty.
func (m *Mut[T]) Set(v T) { m.dirty = true; m.value = v }

// GetAny is a bridge method for ToMap using reflection.
func (m *Mut[T]) GetAny() any { return m.value }

// MarshalJSON implements the json.Marshaler interface.
func (m *Mut[T]) MarshalJSON() ([]byte, error) { return json.Marshal(m.value) }

// UnmarshalJSON implements the json.Unmarshaler interface and marks the value as dirty.
func (m *Mut[T]) UnmarshalJSON(data []byte) error {
	m.dirty = true
	return json.Unmarshal(data, &m.value)
}

// New creates a new Mut instance. If an initial value is provided, it starts as Dirty.
func New[T any](val ...T) Mut[T] {
	if len(val) > 0 {
		return Mut[T]{dirty: true, value: val[0]}
	}
	return Mut[T]{}
}

// ToMap converts structs with Mut fields into map[string]any, including only dirty fields.
func ToMap(obj any) map[string]any {
	// internalMutable is used to identify Mut fields via reflection.
	type internalMutable interface {
		GetAny() any
		Dirty() bool
	}

	out := make(map[string]any)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return out
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Attempt to get the field address to access pointer receiver methods.
		var ptr reflect.Value
		if field.CanAddr() {
			ptr = field.Addr()
		} else {
			// Fallback for structs passed by value.
			c := reflect.New(field.Type())
			c.Elem().Set(field)
			ptr = c
		}

		if m, ok := ptr.Interface().(internalMutable); ok {
			if m.Dirty() {
				tag := t.Field(i).Tag.Get("json")
				key := strings.Split(tag, ",")[0]
				if key == "" || key == "-" {
					key = t.Field(i).Name
				}
				out[key] = m.GetAny()
			}
		}
	}
	return out
}
