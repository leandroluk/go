// oas/builder/request_body_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type RequestBodyBuilder struct {
	requestBody *types.RequestBody
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

func (b *RequestBodyBuilder) ContentJSON(build func(target *MediaTypeBuilder)) *RequestBodyBuilder {
	return b.Content(string(types.ContentType_ApplicationJson), build)
}
