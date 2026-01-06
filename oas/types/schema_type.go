// oas/types/schema_type.go
package types

type SchemaType string

const (
	SchemaType_Object  SchemaType = "object"
	SchemaType_String  SchemaType = "string"
	SchemaType_Integer SchemaType = "integer"
	SchemaType_Number  SchemaType = "number"
	SchemaType_Boolean SchemaType = "boolean"
	SchemaType_Array   SchemaType = "array"
	SchemaType_Null    SchemaType = "null"
)
