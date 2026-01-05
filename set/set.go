package set

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Field represents a value that can be explicitly set via JSON unmarshaling.
// It allows distinguishing between a missing field and a zero-value field.
type Field[T any] struct {
	Value T
	IsSet bool
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It marks the field as set whenever it's present in the JSON input.
func (s *Field[T]) UnmarshalJSON(data []byte) error {
	s.IsSet = true
	return json.Unmarshal(data, &s.Value)
}

// IsSettable defines a common interface for Field types to check their state
// and retrieve values during reflection-based processing.
type IsSettable interface {
	WasSet() bool
	GetValue() any
}

// WasSet returns true if the field was present in the JSON.
func (s Field[T]) WasSet() bool {
	return s.IsSet
}

// GetValue returns the underlying value of the field.
func (s Field[T]) GetValue() any {
	return s.Value
}

// ToMap converts a struct containing Field types into a map[string]any.
// It only includes fields where WasSet() is true. It respects "json" tags
// for key naming.
func ToMap(values any) map[string]any {
	out := make(map[string]any)
	v := reflect.ValueOf(values)

	// Handle pointer to struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return out
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)

		// Check if the field implements IsSettable
		if settable, ok := field.Interface().(IsSettable); ok {
			if settable.WasSet() {
				// Parse JSON tag to get the key name
				tag := structField.Tag.Get("json")
				key := strings.Split(tag, ",")[0]

				if key == "" || key == "-" {
					key = structField.Name
				}
				out[key] = settable.GetValue()
			}
		}
	}
	return out
}
