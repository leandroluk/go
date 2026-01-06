// oas/types/discriminator.go
package types

type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}
