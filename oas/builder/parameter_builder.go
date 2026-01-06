// oas/builder/parameter_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type ParameterBuilder struct {
	parameter *types.Parameter
}

func (b *ParameterBuilder) Name(value string) *ParameterBuilder {
	b.parameter.Name = value
	return b
}

func (b *ParameterBuilder) In(value string) *ParameterBuilder {
	b.parameter.In = value
	return b
}

func (b *ParameterBuilder) Description(value string) *ParameterBuilder {
	b.parameter.Description = value
	return b
}

func (b *ParameterBuilder) Required(value bool) *ParameterBuilder {
	b.parameter.Required = value
	return b
}

func (b *ParameterBuilder) Deprecated(value bool) *ParameterBuilder {
	b.parameter.Deprecated = value
	return b
}

func (b *ParameterBuilder) AllowEmptyValue(value bool) *ParameterBuilder {
	b.parameter.AllowEmptyValue = value
	return b
}

func (b *ParameterBuilder) AllowReserved(value bool) *ParameterBuilder {
	b.parameter.AllowReserved = value
	return b
}

func (b *ParameterBuilder) Example(value any) *ParameterBuilder {
	b.parameter.Example = value
	return b
}

func (b *ParameterBuilder) Schema(build func(target *SchemaBuilder)) *ParameterBuilder {
	schema := &types.Schema{}
	sb := &SchemaBuilder{schema: schema}
	if build != nil {
		build(sb)
	}
	b.parameter.Schema = schema
	return b
}
