// oas/types/server_variable.go
package types

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default"`
	Description string   `json:"description,omitempty"`
}
