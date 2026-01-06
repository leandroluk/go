// oas/builder/schema_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type SchemaBuilder struct {
	schema *types.Schema
}

func (b *SchemaBuilder) Schema() *types.Schema {
	return b.schema
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

func (b *SchemaBuilder) Property(name string, build func(target *SchemaBuilder)) *SchemaBuilder {
	if b.schema.Properties == nil {
		b.schema.Properties = map[string]*types.Schema{}
	}
	child := &types.Schema{}
	cb := &SchemaBuilder{schema: child}
	if build != nil {
		build(cb)
	}
	b.schema.Properties[name] = child
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
