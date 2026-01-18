// oas/builder/schema_builder.go
package builder

import (
	"encoding/json"

	"github.com/leandroluk/go/oas/types"
)

type SchemaBuilder struct {
	schema *types.Schema
}

func (b *SchemaBuilder) Schema() *types.Schema {
	return b.schema
}

// MarshalJSON implementa json.Marshaler
func (b *SchemaBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.schema)
}

// Construtores por tipo

// String cria um builder para schema string
func String() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_String}}}
}

// Integer cria um builder para schema integer
func Integer() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_Integer}}}
}

// Number cria um builder para schema number
func Number() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_Number}}}
}

// Boolean cria um builder para schema boolean
func Boolean() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_Boolean}}}
}

// Array cria um builder para schema array
func Array() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_Array}}}
}

// Object cria um builder para schema object
func Object() *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Type: []types.SchemaType{types.SchemaType_Object}}}
}

// Ref cria um schema com referência
func Ref(ref string) *SchemaBuilder {
	return &SchemaBuilder{schema: &types.Schema{Ref: ref}}
}

func (b *SchemaBuilder) Ref(value string) *SchemaBuilder {
	b.schema.Ref = value
	return b
}

func (b *SchemaBuilder) Type(values ...types.SchemaType) *SchemaBuilder {
	b.schema.Type = append(b.schema.Type, values...)
	return b
}

func (b *SchemaBuilder) Format(value string) *SchemaBuilder {
	b.schema.Format = value
	return b
}

func (b *SchemaBuilder) Title(value string) *SchemaBuilder {
	b.schema.Title = value
	return b
}

func (b *SchemaBuilder) Description(value string) *SchemaBuilder {
	b.schema.Description = value
	return b
}

func (b *SchemaBuilder) Default(value any) *SchemaBuilder {
	b.schema.Default = value
	return b
}

func (b *SchemaBuilder) Const(value any) *SchemaBuilder {
	b.schema.Const = value
	return b
}

func (b *SchemaBuilder) Enum(values ...any) *SchemaBuilder {
	b.schema.Enum = append(b.schema.Enum, values...)
	return b
}

func (b *SchemaBuilder) Property(name string, schema interface{}) *SchemaBuilder {
	if b.schema.Properties == nil {
		b.schema.Properties = make(map[string]*types.Schema)
	}

	switch s := schema.(type) {
	case *SchemaBuilder:
		b.schema.Properties[name] = s.schema
	case *types.Schema:
		b.schema.Properties[name] = s
	}

	return b
}

func (b *SchemaBuilder) Required(names ...string) *SchemaBuilder {
	b.schema.Required = append(b.schema.Required, names...)
	return b
}

func (b *SchemaBuilder) Items(build func(target *SchemaBuilder)) *SchemaBuilder {
	item := &types.Schema{}
	ib := &SchemaBuilder{schema: item}
	if build != nil {
		build(ib)
	}
	b.schema.Items = []*types.Schema{item}
	return b
}

func (b *SchemaBuilder) AllOf(schemaList ...*types.Schema) *SchemaBuilder {
	b.schema.AllOf = append(b.schema.AllOf, schemaList...)
	return b
}

func (b *SchemaBuilder) AnyOf(schemaList ...*types.Schema) *SchemaBuilder {
	b.schema.AnyOf = append(b.schema.AnyOf, schemaList...)
	return b
}

func (b *SchemaBuilder) OneOf(schemaList ...*types.Schema) *SchemaBuilder {
	b.schema.OneOf = append(b.schema.OneOf, schemaList...)
	return b
}

func (b *SchemaBuilder) Not(schema *types.Schema) *SchemaBuilder {
	b.schema.Not = schema
	return b
}

func (b *SchemaBuilder) AdditionalProperties(value any) *SchemaBuilder {
	b.schema.AdditionalProperties = value
	return b
}

func (b *SchemaBuilder) ReadOnly(value bool) *SchemaBuilder {
	b.schema.ReadOnly = &value
	return b
}

func (b *SchemaBuilder) WriteOnly(value bool) *SchemaBuilder {
	b.schema.WriteOnly = &value
	return b
}

func (b *SchemaBuilder) Deprecated(value bool) *SchemaBuilder {
	b.schema.Deprecated = &value
	return b
}

func (b *SchemaBuilder) Example(value any) *SchemaBuilder {
	b.schema.Example = value
	return b
}

// MinLength define comprimento mínimo para strings
func (b *SchemaBuilder) MinLength(n int64) *SchemaBuilder {
	b.schema.MinLength = &n
	return b
}

// MaxLength define comprimento máximo para strings
func (b *SchemaBuilder) MaxLength(n int64) *SchemaBuilder {
	b.schema.MaxLength = &n
	return b
}

// Minimum define valor mínimo para números
func (b *SchemaBuilder) Minimum(n float64) *SchemaBuilder {
	b.schema.Minimum = &n
	return b
}

// Maximum define valor máximo para números
func (b *SchemaBuilder) Maximum(n float64) *SchemaBuilder {
	b.schema.Maximum = &n
	return b
}

// Pattern define expressão regular para validação
func (b *SchemaBuilder) Pattern(p string) *SchemaBuilder {
	b.schema.Pattern = p
	return b
}
