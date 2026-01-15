// search/query.go
package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type FieldConstraint interface{ ~string }

type ProjectMode string

const (
	INCLUDE ProjectMode = "include"
	EXCLUDE ProjectMode = "exclude"
)

type Project[K FieldConstraint] struct {
	Mode   ProjectMode `json:"mode"`
	Fields []K         `json:"fields"`
}

type SortOrder int8

const (
	ASC  SortOrder = 1
	DESC SortOrder = -1
)

type SortItem[K FieldConstraint] struct {
	Field K         `json:"field"`
	Order SortOrder `json:"order"`
}

type Query[TType any, KeyList FieldConstraint] struct {
	Where   *TType              `json:"where,omitempty"`
	Project *Project[KeyList]   `json:"project,omitempty"`
	Sort    []SortItem[KeyList] `json:"sort,omitempty"`
	Limit   *int                `json:"limit,omitempty"`
	Offset  *int                `json:"offset,omitempty"`
}

func (q *Query[TType, KeyList]) UnmarshalJSON(data []byte) error {
	type Alias Query[TType, KeyList]

	aux := &struct {
		*Alias
		Sort json.RawMessage `json:"sort,omitempty"`
	}{Alias: (*Alias)(q)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if len(aux.Sort) == 0 {
		return nil
	}

	first := firstNonSpaceByte(aux.Sort)
	if first == 0 {
		return nil
	}

	if first == '[' {
		var items []SortItem[KeyList]
		if err := json.Unmarshal(aux.Sort, &items); err != nil {
			return err
		}
		q.Sort = items
		return nil
	}

	if first != '{' {
		return fmt.Errorf("invalid sort json: expected object or array")
	}

	dec := json.NewDecoder(bytes.NewReader(aux.Sort))

	tok, err := dec.Token()
	if err != nil {
		return err
	}
	if tok != json.Delim('{') {
		return fmt.Errorf("invalid sort json: expected object")
	}

	q.Sort = nil

	for dec.More() {
		keyToken, err := dec.Token()
		if err != nil {
			return err
		}

		key, ok := keyToken.(string)
		if !ok {
			return fmt.Errorf("invalid sort json: expected string key")
		}

		var order SortOrder
		if err := dec.Decode(&order); err != nil {
			return err
		}

		q.Sort = append(q.Sort, SortItem[KeyList]{
			Field: KeyList(key),
			Order: order,
		})
	}

	return nil
}

func (q *Query[TType, KeyList]) Validate() error {
	return q.validateAgainstType(reflect.TypeFor[TType]())
}

func (q *Query[TType, KeyList]) ValidateAgainst(target any) error {
	if target == nil {
		return fmt.Errorf("target type is nil")
	}
	return q.validateAgainstType(reflect.TypeOf(target))
}

func (q *Query[TType, KeyList]) validateAgainstType(targetType reflect.Type) error {
	targetType = unwrapPointer(targetType)

	if q.Limit != nil && *q.Limit < 0 {
		return fmt.Errorf("invalid limit: %d", *q.Limit)
	}
	if q.Offset != nil && *q.Offset < 0 {
		return fmt.Errorf("invalid offset: %d", *q.Offset)
	}

	if q.Project != nil {
		if q.Project.Mode != INCLUDE && q.Project.Mode != EXCLUDE {
			return fmt.Errorf("invalid project mode: %s", q.Project.Mode)
		}
	}

	jsonFields := collectJSONFields(targetType)

	if q.Project != nil && len(q.Project.Fields) > 0 {
		for _, field := range q.Project.Fields {
			if !jsonFields[string(field)] {
				return fmt.Errorf("invalid projection field: %s", field)
			}
		}
	}

	for _, item := range q.Sort {
		if item.Order != ASC && item.Order != DESC {
			return fmt.Errorf("invalid sort order for field %s: %d", item.Field, item.Order)
		}
		if !jsonFields[string(item.Field)] {
			return fmt.Errorf("invalid sort field: %s", item.Field)
		}
	}

	return nil
}

type Result[TType any] struct {
	Items  []TType `json:"items"`
	Total  int     `json:"total"`
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
}

func firstNonSpaceByte(b []byte) byte {
	for _, c := range b {
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			continue
		}
		return c
	}
	return 0
}

func unwrapPointer(t reflect.Type) reflect.Type {
	for t != nil && t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func collectJSONFields(t reflect.Type) map[string]bool {
	fields := make(map[string]bool)
	collectJSONFieldsInto(fields, t)
	return fields
}

func collectJSONFieldsInto(out map[string]bool, t reflect.Type) {
	t = unwrapPointer(t)
	if t == nil || t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Anonymous {
			collectJSONFieldsInto(out, f.Type)
		}

		tag := f.Tag.Get("json")
		if tag == "" {
			continue
		}

		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}

		out[name] = true
	}
}
