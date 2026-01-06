# oas

Framework-agnostic OpenAPI 3.1 document types + fluent builders for Go.

This module provides:
- OpenAPI 3.1 **static types** under `oas/types`
- A fluent **builder API** under `oas/builder`
- A root **facade** (`package oas`) that re-exports the public surface, so consumers can import only `github.com/leandroluk/go/oas`

## Install

```sh
    go get github.com/leandroluk/go/oas
```

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/leandroluk/go/oas"
)

func main() {
    api := oas.New().
        Title("API").
        Description("API description").
        Version("1.0.0")

    bytes, _ := json.MarshalIndent(api.Document(), "", "  ")
    fmt.Println(string(bytes))
}
```

## Paths and Operations

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/leandroluk/go/oas"
)

func main() {
    api := oas.New().
        Title("API").
        Description("API description").
        Version("1.0.0")

    api.ComponentSchema("User", &oas.Schema{
        Type: []oas.SchemaType{oas.SchemaType_Object},
        Properties: map[string]*oas.Schema{
            "id":   {Type: []oas.SchemaType{oas.SchemaType_String}},
            "name": {Type: []oas.SchemaType{oas.SchemaType_String}},
        },
        Required: []string{"id", "name"},
    })

    api.Path("/users").Get(func(op *oas.OperationBuilder) {
        op.OperationID("ListUsers").
            Summary("List users").
            ResponseRange(2, func(r *oas.ResponseBuilder) {
                r.Description("Success").
                    ContentJSON(func(m *oas.MediaTypeBuilder) {
                        m.Schema(func(s *oas.SchemaBuilder) {
                            s.Type(oas.SchemaType_Array).
                                Items(func(i *oas.SchemaBuilder) {
                                    i.Ref("#/components/schemas/User")
                                })
                        })
                    })
            }).
            Response(http.StatusBadRequest, func(r *oas.ResponseBuilder) {
                r.Description("Bad request")
            }).
            DefaultResponse(func(r *oas.ResponseBuilder) {
                r.Description("Unexpected error")
            })
    })

    bytes, _ := json.MarshalIndent(api.Document(), "", "  ")
    fmt.Println(string(bytes))
}
```

## Response Codes

In OpenAPI, `responses` keys are strings (to support exact codes, ranges, and `default`):
- Exact codes: `"200"`, `"404"`
- Ranges: `"2XX"`, `"4XX"`
- Default: `"default"`

The builder keeps it ergonomic:
- `Response(code int, ...)` -> `"200"`
- `ResponseRange(class int, ...)` -> `"2XX"`
- `DefaultResponse(...)` -> `"default"`

## Package Layout

```
oas/
  oas.go
  types/
    open_api.go
    info.go
    schema.go
    operation.go
    response.go
    ...
  builder/
    new.go
    openapi_builder.go
    path_item_builder.go
    operation_builder.go
    response_builder.go
    request_body_builder.go
    parameter_builder.go
    media_type_builder.go
    schema_builder.go
    content_builder.go
  internal/
    reflection/
    normalize/
```

## Non-Goals

- No HTTP framework bindings (Fiber/Gin/etc.)
- No runtime server/router
- No application-specific metadata dependencies (keep this module reusable)
