// oas/builder/response_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type ResponseBuilder struct {
	response *types.Response
}

func (b *ResponseBuilder) Response() *types.Response {
	return b.response
}

func (b *ResponseBuilder) Description(value string) *ResponseBuilder {
	b.response.Description = value
	return b
}

func (b *ResponseBuilder) Header(name string, build func(target *HeaderBuilder)) *ResponseBuilder {
	if b.response.Headers == nil {
		b.response.Headers = map[string]*types.Header{}
	}
	header := &types.Header{}
	hb := &HeaderBuilder{header: header}
	if build != nil {
		build(hb)
	}
	b.response.Headers[name] = header
	return b
}

func (b *ResponseBuilder) Link(name string, build func(target *LinkBuilder)) *ResponseBuilder {
	if b.response.Links == nil {
		b.response.Links = map[string]*types.Link{}
	}
	link := &types.Link{}
	lb := &LinkBuilder{link: link}
	if build != nil {
		build(lb)
	}
	b.response.Links[name] = link
	return b
}

func (b *ResponseBuilder) Content(contentType string, build func(target *MediaTypeBuilder)) *ResponseBuilder {
	if b.response.Content == nil {
		b.response.Content = map[string]*types.MediaType{}
	}
	media := &types.MediaType{}
	mb := &MediaTypeBuilder{mediaType: media}
	if build != nil {
		build(mb)
	}
	b.response.Content[contentType] = media
	return b
}

func (b *ResponseBuilder) ContentJSON(build func(target *MediaTypeBuilder)) *ResponseBuilder {
	return b.Content(string(types.ContentType_ApplicationJson), build)
}

type HeaderBuilder struct {
	header *types.Header
}

func (b *HeaderBuilder) Description(value string) *HeaderBuilder {
	b.header.Description = value
	return b
}

func (b *HeaderBuilder) Required(value bool) *HeaderBuilder {
	b.header.Required = value
	return b
}

func (b *HeaderBuilder) Deprecated(value bool) *HeaderBuilder {
	b.header.Deprecated = value
	return b
}

func (b *HeaderBuilder) Example(value any) *HeaderBuilder {
	b.header.Example = value
	return b
}

func (b *HeaderBuilder) Schema(build func(target *SchemaBuilder)) *HeaderBuilder {
	schema := &types.Schema{}
	sb := &SchemaBuilder{schema: schema}
	if build != nil {
		build(sb)
	}
	b.header.Schema = schema
	return b
}

type LinkBuilder struct {
	link *types.Link
}

func (b *LinkBuilder) OperationRef(value string) *LinkBuilder {
	b.link.OperationRef = value
	return b
}

func (b *LinkBuilder) OperationID(value string) *LinkBuilder {
	b.link.OperationId = value
	return b
}

func (b *LinkBuilder) Description(value string) *LinkBuilder {
	b.link.Description = value
	return b
}

func (b *LinkBuilder) Parameter(name string, value any) *LinkBuilder {
	if b.link.Parameters == nil {
		b.link.Parameters = map[string]any{}
	}
	b.link.Parameters[name] = value
	return b
}

func (b *LinkBuilder) RequestBody(value any) *LinkBuilder {
	b.link.RequestBody = value
	return b
}

func (b *LinkBuilder) Server(url string, build func(target *types.Server)) *LinkBuilder {
	server := &types.Server{URL: url}
	if build != nil {
		build(server)
	}
	b.link.Server = server
	return b
}
