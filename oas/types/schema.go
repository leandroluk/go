package types

type Schema struct {
	Type                  []SchemaType       `json:"type,omitempty"`
	Format                string             `json:"format,omitempty"`
	Description           string             `json:"description,omitempty"`
	Properties            map[string]*Schema `json:"properties,omitempty"`
	Required              []string           `json:"required,omitempty"`
	Items                 []*Schema          `json:"items,omitempty"`
	Title                 string             `json:"title,omitempty"`
	MultipleOf            *float64           `json:"multipleOf,omitempty"`
	Maximum               *float64           `json:"maximum,omitempty"`
	ExclusiveMaximum      *bool              `json:"exclusiveMaximum,omitempty"`
	Minimum               *float64           `json:"minimum,omitempty"`
	ExclusiveMinimum      *bool              `json:"exclusiveMinimum,omitempty"`
	MaxLength             *int64             `json:"maxLength,omitempty"`
	MinLength             *int64             `json:"minLength,omitempty"`
	Pattern               string             `json:"pattern,omitempty"`
	MaxItems              *int64             `json:"maxItems,omitempty"`
	MinItems              *int64             `json:"minItems,omitempty"`
	UniqueItems           *bool              `json:"uniqueItems,omitempty"`
	MaxProperties         *int64             `json:"maxProperties,omitempty"`
	MinProperties         *int64             `json:"minProperties,omitempty"`
	Enum                  []any              `json:"enum,omitempty"`
	AllOf                 []*Schema          `json:"allOf,omitempty"`
	OneOf                 []*Schema          `json:"oneOf,omitempty"`
	AnyOf                 []*Schema          `json:"anyOf,omitempty"`
	Not                   *Schema            `json:"not,omitempty"`
	AdditionalProperties  any                `json:"additionalProperties,omitempty"`
	Default               any                `json:"default,omitempty"`
	Discriminator         *Discriminator     `json:"discriminator,omitempty"`
	ReadOnly              *bool              `json:"readOnly,omitempty"`
	WriteOnly             *bool              `json:"writeOnly,omitempty"`
	XML                   *XML               `json:"xml,omitempty"`
	ExternalDocs          *ExternalDocs      `json:"externalDocs,omitempty"`
	Example               any                `json:"example,omitempty"`
	Deprecated            *bool              `json:"deprecated,omitempty"`
	DependentSchemas      map[string]*Schema `json:"dependentSchemas,omitempty"`
	UnevaluatedItems      any                `json:"unevaluatedItems,omitempty"`
	UnevaluatedProperties any                `json:"unevaluatedProperties,omitempty"`
	If                    *Schema            `json:"if,omitempty"`
	Then                  *Schema            `json:"then,omitempty"`
	Else                  *Schema            `json:"else,omitempty"`
	ContentMediaType      string             `json:"contentMediaType,omitempty"`
	ContentEncoding       string             `json:"contentEncoding,omitempty"`
	Ref                   string             `json:"$ref,omitempty"`
	Const                 any                `json:"const,omitempty"`
}
