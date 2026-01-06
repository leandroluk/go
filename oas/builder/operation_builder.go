// oas/builder/operation_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type OperationBuilder struct {
	operation *types.Operation
}

func (b *OperationBuilder) Operation() *types.Operation {
	return b.operation
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
