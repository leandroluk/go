// oas/types/operation.go
package types

type PathOperation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	OperationId  string                 `json:"operationId,omitempty"`
	Parameters   []*Parameter           `json:"parameters,omitempty"`
	RequestBody  *RequestBody           `json:"requestBody,omitempty"`
	Responses    map[string]*Response   `json:"responses,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty"`
	ExternalDocs *ExternalDocs          `json:"externalDocs,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty"`
	Servers      []*Server              `json:"servers,omitempty"`
}
