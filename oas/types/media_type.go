// oas/types/media_type.go
package types

type MediaType struct {
	Schema   *Schema                   `json:"schema,omitempty"`
	Example  any                       `json:"example,omitempty"`
	Examples map[string]*ExampleObject `json:"examples,omitempty"`
	Encoding map[string]*Encoding      `json:"encoding,omitempty"`
}
