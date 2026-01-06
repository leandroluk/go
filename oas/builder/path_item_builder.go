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

func (b *PathItemBuilder) Get(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Get = operation
	return b
}

func (b *PathItemBuilder) Post(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Post = operation
	return b
}

func (b *PathItemBuilder) Put(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Put = operation
	return b
}

func (b *PathItemBuilder) Delete(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Delete = operation
	return b
}

func (b *PathItemBuilder) Patch(build func(target *OperationBuilder)) *PathItemBuilder {
	operation := newOperation()
	ob := &OperationBuilder{operation: operation}
	if build != nil {
		build(ob)
	}
	b.item.Patch = operation
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

func newOperation() *types.Operation {
	return &types.Operation{
		Responses: map[string]*types.Response{},
	}
}
