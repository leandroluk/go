// oas/builder/openapi_builder.go
package builder

import (
	"encoding/json"

	"github.com/leandroluk/go/oas/types"
)

type OpenAPIBuilder struct {
	document *types.OpenAPI
}

func (b *OpenAPIBuilder) Document() *types.OpenAPI {
	return b.document
}

// MarshalJSON implementa json.Marshaler
func (b *OpenAPIBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.document)
}

func (b *OpenAPIBuilder) Title(value string) *OpenAPIBuilder {
	b.ensureInfo()
	b.document.Info.Title = value
	return b
}

func (b *OpenAPIBuilder) Description(value string) *OpenAPIBuilder {
	b.ensureInfo()
	b.document.Info.Description = value
	return b
}

func (b *OpenAPIBuilder) Version(value string) *OpenAPIBuilder {
	b.ensureInfo()
	b.document.Info.Version = value
	return b
}

func (b *OpenAPIBuilder) Summary(value string) *OpenAPIBuilder {
	b.ensureInfo()
	b.document.Info.Summary = value
	return b
}

func (b *OpenAPIBuilder) TermsOfService(value string) *OpenAPIBuilder {
	b.ensureInfo()
	b.document.Info.TermsOfService = value
	return b
}

func (b *OpenAPIBuilder) Contact(build func(target *types.Contact)) *OpenAPIBuilder {
	b.ensureInfo()
	if b.document.Info.Contact == nil {
		b.document.Info.Contact = &types.Contact{}
	}
	if build != nil {
		build(b.document.Info.Contact)
	}
	return b
}

func (b *OpenAPIBuilder) License(build func(target *types.License)) *OpenAPIBuilder {
	b.ensureInfo()
	if b.document.Info.License == nil {
		b.document.Info.License = &types.License{Name: ""}
	}
	if build != nil {
		build(b.document.Info.License)
	}
	return b
}

func (b *OpenAPIBuilder) Server(url string, build func(target *types.Server)) *OpenAPIBuilder {
	server := &types.Server{URL: url}
	if build != nil {
		build(server)
	}
	b.document.Servers = append(b.document.Servers, server)
	return b
}

func (b *OpenAPIBuilder) Tag(name string, build func(target *types.Tag)) *OpenAPIBuilder {
	tag := &types.Tag{Name: name}
	if build != nil {
		build(tag)
	}
	b.document.Tags = append(b.document.Tags, tag)
	return b
}

func (b *OpenAPIBuilder) Security(requirement types.SecurityRequirement) *OpenAPIBuilder {
	b.document.Security = append(b.document.Security, &requirement)
	return b
}

func (b *OpenAPIBuilder) ExternalDocs(build func(target *types.ExternalDocs)) *OpenAPIBuilder {
	if b.document.ExternalDocs == nil {
		b.document.ExternalDocs = &types.ExternalDocs{URL: ""}
	}
	if build != nil {
		build(b.document.ExternalDocs)
	}
	return b
}

func (b *OpenAPIBuilder) Path(path string) *PathItemBuilder {
	b.ensurePaths()

	key := normalizePath(path)
	item := b.document.Paths[key]
	if item == nil {
		item = &types.PathItem{}
		b.document.Paths[key] = item
	}

	return &PathItemBuilder{
		openapi: b,
		path:    key,
		item:    item,
	}
}

func (b *OpenAPIBuilder) Webhook(name string) *PathItemBuilder {
	if b.document.Webhooks == nil {
		b.document.Webhooks = types.Webhooks{}
	}
	item := b.document.Webhooks[name]
	if item == nil {
		item = &types.PathItem{}
		b.document.Webhooks[name] = item
	}

	return &PathItemBuilder{
		openapi: b,
		path:    name,
		item:    item,
	}
}

func (b *OpenAPIBuilder) ComponentSchema(name string, schema *types.Schema) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.Schemas == nil {
		b.document.Components.Schemas = map[string]*types.Schema{}
	}
	b.document.Components.Schemas[name] = schema
	return b
}

func (b *OpenAPIBuilder) ComponentResponse(name string, response *types.Response) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.Responses == nil {
		b.document.Components.Responses = map[string]*types.Response{}
	}
	b.document.Components.Responses[name] = response
	return b
}

func (b *OpenAPIBuilder) ComponentParameter(name string, parameter *types.Parameter) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.Parameters == nil {
		b.document.Components.Parameters = map[string]*types.Parameter{}
	}
	b.document.Components.Parameters[name] = parameter
	return b
}

func (b *OpenAPIBuilder) ComponentRequestBody(name string, requestBody *types.RequestBody) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.RequestBodies == nil {
		b.document.Components.RequestBodies = map[string]*types.RequestBody{}
	}
	b.document.Components.RequestBodies[name] = requestBody
	return b
}

func (b *OpenAPIBuilder) ComponentHeader(name string, header *types.Header) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.Headers == nil {
		b.document.Components.Headers = map[string]*types.Header{}
	}
	b.document.Components.Headers[name] = header
	return b
}

func (b *OpenAPIBuilder) ComponentSecurityScheme(name string, scheme *types.SecurityScheme) *OpenAPIBuilder {
	b.ensureComponents()
	if b.document.Components.SecuritySchemes == nil {
		b.document.Components.SecuritySchemes = map[string]*types.SecurityScheme{}
	}
	b.document.Components.SecuritySchemes[name] = scheme
	return b
}

func (b *OpenAPIBuilder) ensureInfo() {
	if b.document.Info == nil {
		b.document.Info = &types.Info{
			Title:   "",
			Version: "0.0.0",
		}
	}
}

func (b *OpenAPIBuilder) ensurePaths() {
	if b.document.Paths == nil {
		b.document.Paths = types.Paths{}
	}
}

func (b *OpenAPIBuilder) ensureComponents() {
	if b.document.Components == nil {
		b.document.Components = &types.Components{}
	}
}
