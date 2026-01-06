// oas/types/open_api.go
package types

type OpenAPI struct {
	OpenAPI           string                 `json:"openapi"`
	Info              *Info                  `json:"info"`
	JsonSchemaDialect string                 `json:"jsonSchemaDialect,omitempty"`
	Servers           []*Server              `json:"servers,omitempty"`
	Paths             Paths                  `json:"paths,omitempty"`
	Webhooks          Webhooks               `json:"webhooks,omitempty"`
	Components        *Components            `json:"components,omitempty"`
	Security          []*SecurityRequirement `json:"security,omitempty"`
	Tags              []*Tag                 `json:"tags,omitempty"`
	ExternalDocs      *ExternalDocs          `json:"externalDocs,omitempty"`
}

func (o *OpenAPI) AddPathItem(path string) *PathItem {
	if o.Paths == nil {
		o.Paths = make(Paths)
	}

	pathItem := o.Paths[path]
	if pathItem == nil {
		pathItem = &PathItem{}
		o.Paths[path] = pathItem
	}
	return pathItem
}
