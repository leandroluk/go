// oas/types/header.go
package types

type Header struct {
	Description     string  `json:"description,omitempty"`
	Required        bool    `json:"required,omitempty"`
	Schema          *Schema `json:"schema,omitempty"`
	Example         any     `json:"example,omitempty"`
	Deprecated      bool    `json:"deprecated,omitempty"`
	AllowEmptyValue bool    `json:"allowEmptyValue,omitempty"`
}
