// oas/types/tag.go
package types

type Tag struct {
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}
