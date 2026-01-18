// oas/types/path_item.go
package types

type PathItem struct {
	Ref         string         `json:"$ref,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Get         *PathOperation `json:"get,omitempty"`
	Post        *PathOperation `json:"post,omitempty"`
	Put         *PathOperation `json:"put,omitempty"`
	Delete      *PathOperation `json:"delete,omitempty"`
	Options     *PathOperation `json:"options,omitempty"`
	Head        *PathOperation `json:"head,omitempty"`
	Patch       *PathOperation `json:"patch,omitempty"`
	Trace       *PathOperation `json:"trace,omitempty"`
	Servers     []*Server      `json:"servers,omitempty"`
	Parameters  []*Parameter   `json:"parameters,omitempty"`
}
