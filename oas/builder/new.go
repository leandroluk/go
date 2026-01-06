// oas/builder/new.go
package builder

import "github.com/leandroluk/go/oas/types"

func New() *OpenAPIBuilder {
	document := &types.OpenAPI{
		OpenAPI: "3.1.0",
		Info: &types.Info{
			Title:   "",
			Version: "0.0.0",
		},
		Paths:      types.Paths{},
		Components: &types.Components{},
	}

	return &OpenAPIBuilder{document: document}
}
