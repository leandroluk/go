// oas/builder/request_body_builder.go
package builder

import (
	"encoding/json"

	"github.com/leandroluk/go/oas/types"
)

type RequestBodyBuilder struct {
	requestBody *types.RequestBody
}

// Body cria um builder fluente para request body
func Body() *RequestBodyBuilder {
	return &RequestBodyBuilder{
		requestBody: &types.RequestBody{},
	}
}

// MarshalJSON implementa json.Marshaler
func (b *RequestBodyBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.requestBody)
}

func (b *RequestBodyBuilder) Description(value string) *RequestBodyBuilder {
	b.requestBody.Description = value
	return b
}

func (b *RequestBodyBuilder) Required(value bool) *RequestBodyBuilder {
	b.requestBody.Required = value
	return b
}

func (b *RequestBodyBuilder) Content(contentType string, build func(target *MediaTypeBuilder)) *RequestBodyBuilder {
	if b.requestBody.Content == nil {
		b.requestBody.Content = map[string]*types.MediaType{}
	}
	media := &types.MediaType{}
	mb := &MediaTypeBuilder{mediaType: media}
	if build != nil {
		build(mb)
	}
	b.requestBody.Content[contentType] = media
	return b
}

func (b *RequestBodyBuilder) ContentJSON(schemaOrCallback interface{}) *RequestBodyBuilder {
	if b.requestBody.Content == nil {
		b.requestBody.Content = map[string]*types.MediaType{}
	}

	// Suporta tanto schema direto quanto callback (retrocompat)
	switch v := schemaOrCallback.(type) {
	case func(*MediaTypeBuilder):
		return b.Content(string(types.ContentType_ApplicationJson), v)
	case *SchemaBuilder:
		b.requestBody.Content[string(types.ContentType_ApplicationJson)] = &types.MediaType{Schema: v.schema}
	case *types.Schema:
		b.requestBody.Content[string(types.ContentType_ApplicationJson)] = &types.MediaType{Schema: v}
	}
	return b
}
