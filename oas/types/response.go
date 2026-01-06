// oas/types/response.go
package types

type Response struct {
	Description string                `json:"description"`
	Content     map[string]*MediaType `json:"content,omitempty"`
	Headers     map[string]*Header    `json:"headers,omitempty"`
	Links       map[string]*Link      `json:"links,omitempty"`
}
