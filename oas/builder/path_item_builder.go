// oas/builder/path_item_builder.go
package builder

import "github.com/leandroluk/go/oas/types"

type PathItemBuilder struct {
	openapi *OpenAPIBuilder
	path    string
	item    *types.PathItem
}

func (b *PathItemBuilder) Ref(value string) *PathItemBuilder {
	b.item.Ref = value
	return b
}

func (b *PathItemBuilder) Summary(value string) *PathItemBuilder {
	b.item.Summary = value
	return b
}

func (b *PathItemBuilder) Description(value string) *PathItemBuilder {
	b.item.Description = value
	return b
}

func (b *PathItemBuilder) Server(url string, build func(target *types.Server)) *PathItemBuilder {
	server := &types.Server{URL: url}
	if build != nil {
		build(server)
	}
	b.item.Servers = append(b.item.Servers, server)
	return b
}

func (b *PathItemBuilder) Parameter(build func(target *ParameterBuilder)) *PathItemBuilder {
	parameter := &types.Parameter{}
	builder := &ParameterBuilder{parameter: parameter}
	if build != nil {
		build(builder)
	}
	b.item.Parameters = append(b.item.Parameters, parameter)
	return b
}

func (b *PathItemBuilder) Get(op *OperationBuilder) *PathItemBuilder {
	if op != nil {
		b.item.Get = op.operation
	}
	return b
}

func (b *PathItemBuilder) Post(op *OperationBuilder) *PathItemBuilder {
	if op != nil {
		b.item.Post = op.operation
	}
	return b
}

func (b *PathItemBuilder) Put(op *OperationBuilder) *PathItemBuilder {
	if op != nil {
		b.item.Put = op.operation
	}
	return b
}

func (b *PathItemBuilder) Delete(op *OperationBuilder) *PathItemBuilder {
	if op != nil {
		b.item.Delete = op.operation
	}
	return b
}

func (b *PathItemBuilder) Patch(op *OperationBuilder) *PathItemBuilder {
	if op != nil {
		b.item.Patch = op.operation
	}
	return b
}

func (b *PathItemBuilder) Head(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Head = operation
	return b
}

func (b *PathItemBuilder) Options(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Options = operation
	return b
}

func (b *PathItemBuilder) Trace(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Trace = operation
	return b
}

func newOperation() *types.PathOperation {
	return &types.PathOperation{
		Responses: map[string]*types.Response{},
	}
}
