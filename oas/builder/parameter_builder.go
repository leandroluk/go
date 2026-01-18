// oas/builder/parameter_builder.go
package builder

import (
	"encoding/json"

	"github.com/leandroluk/go/oas/types"
)

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

// Construtores

// InPath cria um parameter de path
func InPath(name string, schema interface{}) *ParameterBuilder {
	return createParam(name, "path", schema).Required(true)
}

// InQuery cria um parameter de query
func InQuery(name string, schema interface{}) *ParameterBuilder {
	return createParam(name, "query", schema)
}

// InHeader cria um parameter de header
func InHeader(name string, schema interface{}) *ParameterBuilder {
	return createParam(name, "header", schema)
}

// InCookie cria um parameter de cookie
func InCookie(name string, schema interface{}) *ParameterBuilder {
	return createParam(name, "cookie", schema)
}

func createParam(name, in string, schema interface{}) *ParameterBuilder {
	var s *types.Schema
	switch v := schema.(type) {
	case *SchemaBuilder:
		s = v.schema
	case *types.Schema:
		s = v
	}

	return &ParameterBuilder{
		parameter: &types.Parameter{
			Name:   name,
			In:     in,
			Schema: s,
		},
	}
}

// MarshalJSON implementa json.Marshaler
func (b *ParameterBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.parameter)
}
