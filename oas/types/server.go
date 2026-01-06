// oas/types/server.go
package types

type Server struct {
	URL         string                     `json:"url"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}
