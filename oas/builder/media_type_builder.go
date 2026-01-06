// oas/builder/media_type_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type MediaTypeBuilder struct {
	mediaType *types.MediaType
}

func (b *MediaTypeBuilder) Schema(build func(target *SchemaBuilder)) *MediaTypeBuilder {
	schema := &types.Schema{}
	sb := &SchemaBuilder{schema: schema}
	if build != nil {
		build(sb)
	}
	b.mediaType.Schema = schema
	return b
}

func (b *MediaTypeBuilder) SchemaRef(ref string) *MediaTypeBuilder {
	b.mediaType.Schema = &types.Schema{Ref: ref}
	return b
}

func (b *MediaTypeBuilder) Example(value any) *MediaTypeBuilder {
	b.mediaType.Example = value
	return b
}

func (b *MediaTypeBuilder) ExampleNamed(name string, build func(target *ExampleObjectBuilder)) *MediaTypeBuilder {
	if b.mediaType.Examples == nil {
		b.mediaType.Examples = map[string]*types.ExampleObject{}
	}
	ex := &types.ExampleObject{}
	eb := &ExampleObjectBuilder{example: ex}
	if build != nil {
		build(eb)
	}
	b.mediaType.Examples[name] = ex
	return b
}

func (b *MediaTypeBuilder) Encoding(name string, build func(target *EncodingBuilder)) *MediaTypeBuilder {
	if b.mediaType.Encoding == nil {
		b.mediaType.Encoding = map[string]*types.Encoding{}
	}
	enc := &types.Encoding{}
	eb := &EncodingBuilder{encoding: enc}
	if build != nil {
		build(eb)
	}
	b.mediaType.Encoding[name] = enc
	return b
}

type ExampleObjectBuilder struct {
	example *types.ExampleObject
}

func (b *ExampleObjectBuilder) Summary(value string) *ExampleObjectBuilder {
	b.example.Summary = value
	return b
}

func (b *ExampleObjectBuilder) Description(value string) *ExampleObjectBuilder {
	b.example.Description = value
	return b
}

func (b *ExampleObjectBuilder) Value(value any) *ExampleObjectBuilder {
	b.example.Value = value
	return b
}

func (b *ExampleObjectBuilder) ExternalValue(value string) *ExampleObjectBuilder {
	b.example.ExternalValue = value
	return b
}

type EncodingBuilder struct {
	encoding *types.Encoding
}

func (b *EncodingBuilder) ContentType(value string) *EncodingBuilder {
	b.encoding.ContentType = value
	return b
}

func (b *EncodingBuilder) Style(value string) *EncodingBuilder {
	b.encoding.Style = value
	return b
}

func (b *EncodingBuilder) Explode(value bool) *EncodingBuilder {
	b.encoding.Explode = &value
	return b
}

func (b *EncodingBuilder) AllowReserved(value bool) *EncodingBuilder {
	b.encoding.AllowReserved = &value
	return b
}

func (b *EncodingBuilder) Header(name string, build func(target *HeaderBuilder)) *EncodingBuilder {
	if b.encoding.Headers == nil {
		b.encoding.Headers = map[string]*types.Header{}
	}
	header := &types.Header{}
	hb := &HeaderBuilder{header: header}
	if build != nil {
		build(hb)
	}
	b.encoding.Headers[name] = header
	return b
}
