// oas/types/parameter.go
package types

type Parameter struct {
	Name            string  `json:"name"`
	In              string  `json:"in"`
	Description     string  `json:"description,omitempty"`
	Required        bool    `json:"required,omitempty"`
	Schema          *Schema `json:"schema,omitempty"`
	Example         any     `json:"example,omitempty"`
	Deprecated      bool    `json:"deprecated,omitempty"`
	AllowEmptyValue bool    `json:"allowEmptyValue,omitempty"`
	AllowReserved   bool    `json:"allowReserved,omitempty"`
}
