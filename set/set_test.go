package set

import (
	"encoding/json"
	"testing"
)

type UserUpdate struct {
	Name  Field[string] `json:"name"`
	Age   Field[int]    `json:"age"`
	Email Field[string] `json:"email"`
}

func TestField_Unmarshal(t *testing.T) {
	t.Run("Partial JSON input", func(t *testing.T) {
		jsonData := []byte(`{"name": "Leandro"}`)
		var update UserUpdate

		err := json.Unmarshal(jsonData, &update)
		if err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if !update.Name.IsSet {
			t.Error("Expected Name to be set")
		}
		if update.Name.Value != "Leandro" {
			t.Errorf("Expected Name 'Leandro', got %s", update.Name.Value)
		}
		if update.Age.IsSet {
			t.Error("Age should not be set")
		}
	})

	t.Run("Explicit zero value", func(t *testing.T) {
		jsonData := []byte(`{"age": 0}`)
		var update UserUpdate

		json.Unmarshal(jsonData, &update)

		if !update.Age.IsSet {
			t.Error("Expected Age to be set even if value is 0")
		}
		if update.Age.Value != 0 {
			t.Errorf("Expected Age 0, got %d", update.Age.Value)
		}
	})
}

func TestToMap(t *testing.T) {
	t.Run("Generate map with only set fields", func(t *testing.T) {
		update := UserUpdate{
			Name: Field[string]{Value: "Luk", IsSet: true},
			// Age is omitted (IsSet = false)
			Email: Field[string]{Value: "test@test.com", IsSet: true},
		}

		resultMap := ToMap(update)

		if len(resultMap) != 2 {
			t.Errorf("Expected map size 2, got %d", len(resultMap))
		}

		if _, ok := resultMap["name"]; !ok {
			t.Error("Key 'name' should exist in map")
		}

		if _, ok := resultMap["age"]; ok {
			t.Error("Key 'age' should not exist in map")
		}
	})

	t.Run("Handle pointer to struct", func(t *testing.T) {
		update := &UserUpdate{
			Name: Field[string]{Value: "Pointer Test", IsSet: true},
		}

		resultMap := ToMap(update)
		if resultMap["name"] != "Pointer Test" {
			t.Errorf("Expected 'Pointer Test', got %v", resultMap["name"])
		}
	})
}
