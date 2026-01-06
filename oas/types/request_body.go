// oas/types/request_body.go
package types

type RequestBody struct {
	Description string                `json:"description,omitempty"`
	Content     map[string]*MediaType `json:"content"`
	Required    bool                  `json:"required,omitempty"`
}
