// oas/types/path_item.go
package types

type PathItem struct {
	Ref         string       `json:"$ref,omitempty"`
	Summary     string       `json:"summary,omitempty"`
	Description string       `json:"description,omitempty"`
	Get         *Operation   `json:"get,omitempty"`
	Post        *Operation   `json:"post,omitempty"`
	Put         *Operation   `json:"put,omitempty"`
	Delete      *Operation   `json:"delete,omitempty"`
	Options     *Operation   `json:"options,omitempty"`
	Head        *Operation   `json:"head,omitempty"`
	Patch       *Operation   `json:"patch,omitempty"`
	Trace       *Operation   `json:"trace,omitempty"`
	Servers     []*Server    `json:"servers,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
}
