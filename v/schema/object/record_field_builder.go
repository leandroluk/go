package object

import (
	"reflect"

	"github.com/leandroluk/go/v/internal/ast"
	"github.com/leandroluk/go/v/internal/engine"
	"github.com/leandroluk/go/v/internal/issues"
	"github.com/leandroluk/go/v/schema"
)

type RecordFieldBuilder[T any] struct {
	schema    *Schema[T]
	fieldInfo fieldInfo[T]

	required bool
	min      *int
	max      *int
	len      *int

	keySchema   schema.AnySchema
	valueSchema schema.AnySchema

	fieldIndex int
}

func (b *RecordFieldBuilder[T]) Required() *RecordFieldBuilder[T] {
	b.required = true
	return b.build()
}

func (b *RecordFieldBuilder[T]) Min(min int) *RecordFieldBuilder[T] {
	b.min = &min
	return b.build()
}

func (b *RecordFieldBuilder[T]) Max(max int) *RecordFieldBuilder[T] {
	b.max = &max
	return b.build()
}

func (b *RecordFieldBuilder[T]) Len(len int) *RecordFieldBuilder[T] {
	b.len = &len
	return b.build()
}

func (b *RecordFieldBuilder[T]) build() *RecordFieldBuilder[T] {
	mapType := b.fieldInfo.fieldType
	if mapType.Kind() != reflect.Map {
		return b
	}

	validator := func(context *engine.Context, value any) (any, error) {
		val, ok := value.(ast.Value)
		if !ok {
			return nil, nil
		}

		if val.IsMissing() {
			if b.required {
				context.AddIssue("object.required", "required")
				return nil, nil
			}
			return nil, nil
		}
		if val.IsNull() {
			return nil, nil
		}

		if val.Kind != ast.KindObject {
			context.AddIssue("record.type", "expected object/map")
			return nil, nil
		}

		obj := val.Object // map[string]Value

		count := len(obj)

		if b.min != nil && count < *b.min {
			context.AddIssueWithMeta("record.min", "too few items", map[string]any{"min": *b.min, "actual": count})
			return nil, nil
		}
		if b.max != nil && count > *b.max {
			context.AddIssueWithMeta("record.max", "too many items", map[string]any{"max": *b.max, "actual": count})
			return nil, nil
		}
		if b.len != nil && count != *b.len {
			context.AddIssueWithMeta("record.len", "invalid length", map[string]any{"len": *b.len, "actual": count})
			return nil, nil
		}

		resultMap := reflect.MakeMapWithSize(mapType, count)
		basePath := context.PathString()

		for key, item := range obj {
			// Validate Key
			if b.keySchema != nil {
				keyVal := ast.StringValue(key)
				_, err := b.keySchema.ValidateAny(keyVal, context.Options)
				if err != nil {
					if vErr, ok := err.(issues.ValidationError); ok {
						for _, issue := range vErr.Issues {
							context.AddIssueWithMeta("record.key", "invalid key", map[string]any{"key": key, "details": issue.Message})
						}
					}
				}
			}

			// Validate Value
			if b.valueSchema != nil {
				itemRes, err := b.valueSchema.ValidateAny(item, context.Options)
				if err != nil {
					if vErr, ok := err.(issues.ValidationError); ok {
						for _, issue := range vErr.Issues {
							// Construct relative path for this item
							// If key is "myKey", path is "myKey" or "myKey.subPath"
							// We use bracket style if it helps, but JSON pointer uses /myKey/subPath
							// V uses standard dot notation usually.

							var itemRelPath string
							if issue.Path != "" {
								if issue.Path[0] == '[' {
									itemRelPath = key + issue.Path
								} else {
									itemRelPath = key + "." + issue.Path
								}
							} else {
								itemRelPath = key
							}

							var fullPath string
							if basePath != "" {
								fullPath = basePath + "." + itemRelPath
							} else {
								fullPath = itemRelPath
							}

							issue.Path = fullPath
							context.Issues.Add(issue)
						}
					} else {
						return nil, err
					}
				}
				if itemRes != nil {
					keyType := mapType.Key()
					var keyVal reflect.Value
					if keyType.Kind() == reflect.String {
						keyVal = reflect.ValueOf(key)
					} else {
						// Simplistic coercion fallback/skip
						continue
					}

					resultMap.SetMapIndex(keyVal, reflect.ValueOf(itemRes))
				}
			}
		}

		return resultMap.Interface(), nil
	}

	compiled, err := newFieldFromInfo(b.fieldInfo, validator)
	if err != nil {
		b.schema.buildError = err
		return b
	}
	compiled.required = b.required

	if b.fieldIndex == -1 {
		b.schema.fields = append(b.schema.fields, compiled)
		b.schema.lastFieldIndex = len(b.schema.fields) - 1
		b.fieldIndex = b.schema.lastFieldIndex
	} else {
		b.schema.fields[b.fieldIndex] = compiled
		b.schema.lastFieldIndex = b.fieldIndex
	}

	return b
}
