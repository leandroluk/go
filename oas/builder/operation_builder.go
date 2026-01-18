// oas/builder/operation_builder.go
package builder

import (
	"fmt"
	"encoding/json"

	"github.com/leandroluk/go/oas/types"
)

type OperationBuilder struct {
	operation *types.PathOperation
}

func (b *OperationBuilder) Operation() *types.PathOperation {
	return b.operation
}

// MarshalJSON implementa json.Marshaler
func (b *OperationBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.operation)
}

// Operation cria um builder fluente para operação
func Operation(operationId string) *OperationBuilder {
	return &OperationBuilder{
		operation: &types.PathOperation{
			OperationId: operationId,
		},
	}
}

func (b *OperationBuilder) Tag(value string) *OperationBuilder {
	b.operation.Tags = append(b.operation.Tags, value)
	return b
}

func (b *OperationBuilder) Tags(values ...string) *OperationBuilder {
	b.operation.Tags = append(b.operation.Tags, values...)
	return b
}

func (b *OperationBuilder) Summary(value string) *OperationBuilder {
	b.operation.Summary = value
	return b
}

func (b *OperationBuilder) Description(value string) *OperationBuilder {
	b.operation.Description = value
	return b
}

func (b *OperationBuilder) OperationID(value string) *OperationBuilder {
	b.operation.OperationId = value
	return b
}

func (b *OperationBuilder) Deprecated(value bool) *OperationBuilder {
	b.operation.Deprecated = value
	return b
}

func (b *OperationBuilder) ExternalDocs(build func(target *types.ExternalDocs)) *OperationBuilder {
	if b.operation.ExternalDocs == nil {
		b.operation.ExternalDocs = &types.ExternalDocs{URL: ""}
	}
	if build != nil {
		build(b.operation.ExternalDocs)
	}
	return b
}

func (b *OperationBuilder) Security(requirement types.SecurityRequirement) *OperationBuilder {
	b.operation.Security = append(b.operation.Security, &requirement)
	return b
}

func (b *OperationBuilder) Server(url string, build func(target *types.Server)) *OperationBuilder {
	server := &types.Server{URL: url}
	if build != nil {
		build(server)
	}
	b.operation.Servers = append(b.operation.Servers, server)
	return b
}

func (b *OperationBuilder) Parameter(build func(target *ParameterBuilder)) *OperationBuilder {
	parameter := &types.Parameter{}
	builder := &ParameterBuilder{parameter: parameter}
	if build != nil {
		build(builder)
	}
	b.operation.Parameters = append(b.operation.Parameters, parameter)
	return b
}

func (b *OperationBuilder) RequestBody(build func(target *RequestBodyBuilder)) *OperationBuilder {
	requestBody := &types.RequestBody{
		Content: map[string]*types.MediaType{},
	}
	rb := &RequestBodyBuilder{requestBody: requestBody}
	if build != nil {
		build(rb)
	}
	b.operation.RequestBody = requestBody
	return b
}

func (b *OperationBuilder) Response(code int, build func(target *ResponseBuilder)) *OperationBuilder {
	return b.ResponseKey(statusCodeKey(code), build)
}

func (b *OperationBuilder) ResponseRange(class int, build func(target *ResponseBuilder)) *OperationBuilder {
	return b.ResponseKey(statusRangeKey(class), build)
}

func (b *OperationBuilder) DefaultResponse(build func(target *ResponseBuilder)) *OperationBuilder {
	return b.ResponseKey(defaultResponseKey(), build)
}

func (b *OperationBuilder) ResponseKey(key string, build func(target *ResponseBuilder)) *OperationBuilder {
	if b.operation.Responses == nil {
		b.operation.Responses = map[string]*types.Response{}
	}

	response := &types.Response{Description: ""}
	rb := &ResponseBuilder{response: response}
	if build != nil {
		build(rb)
	}
	b.operation.Responses[key] = response
	return b
}

// Parameters adiciona múltiplos parâmetros de uma vez
func (b *OperationBuilder) Parameters(params ...*ParameterBuilder) *OperationBuilder {
	for _, p := range params {
		b.operation.Parameters = append(b.operation.Parameters, p.parameter)
	}
	return b
}

// Responses adiciona múltiplas respostas de uma vez
func (b *OperationBuilder) Responses(responses ...*ResponseWithCode) *OperationBuilder {
	if b.operation.Responses == nil {
		b.operation.Responses = make(map[string]*types.Response)
	}
	for _, r := range responses {
		b.operation.Responses[r.code] = r.response.response
	}
	return b
}

// ResponseWithCode agrupa código e response
type ResponseWithCode struct {
	code     string
	response *ResponseBuilder
}

// ResponseCode cria response com código específico
func ResponseCode(code int) *ResponseWithCode {
	return &ResponseWithCode{
		code:     fmt.Sprintf("%d", code),
		response: &ResponseBuilder{response: &types.Response{}},
	}
}

// ResponseRange cria response com range (2XX, 4XX, etc.)
func ResponseRange(class int) *ResponseWithCode {
	return &ResponseWithCode{
		code:     fmt.Sprintf("%dXX", class),
		response: &ResponseBuilder{response: &types.Response{}},
	}
}

// ResponseDefault cria response default
func ResponseDefault() *ResponseWithCode {
	return &ResponseWithCode{
		code:     "default",
		response: &ResponseBuilder{response: &types.Response{}},
	}
}

// Description adiciona descrição ao response
func (r *ResponseWithCode) Description(desc string) *ResponseWithCode {
	r.response.response.Description = desc
	return r
}

// ContentJSON adiciona content JSON ao response
func (r *ResponseWithCode) ContentJSON(schema interface{}) *ResponseWithCode {
	if r.response.response.Content == nil {
		r.response.response.Content = make(map[string]*types.MediaType)
	}

	var s *types.Schema
	switch v := schema.(type) {
	case *SchemaBuilder:
		s = v.schema
	case *types.Schema:
		s = v
	}

	r.response.response.Content[string(types.ContentType_ApplicationJson)] = &types.MediaType{
		Schema: s,
	}
	return r
}

// MarshalJSON para ResponseWithCode
func (r *ResponseWithCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.response.response)
}

