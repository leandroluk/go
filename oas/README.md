# oas

Type-safe OpenAPI 3.1 document builder for Go with 100% fluent API.

## Installation

```sh
go get github.com/leandroluk/go/oas
```

## Quick Start

```go
package main

import (
    "encoding/json"
    "os"
    "github.com/leandroluk/go/oas"
)

func main() {
    api := oas.New().
        Title("My API").
        Version("1.0.0")

    api.ComponentSchema("User",
        oas.Object().
            Required("id", "name").
            Property("id", oas.String().Example("usr_123")).
            Property("name", oas.String().MinLength(1)),
    )

    api.Path("/users").Get(
        oas.Operation("listUsers").
            Summary("List users").
            Responses(
                oas.ResponseCode(200).
                    Description("Success").
                    ContentJSON(
                        oas.Array().Items(oas.Ref("#/components/schemas/User"))
                    ),
            ),
    )

    json.NewEncoder(os.Stdout).Encode(api)
}
```

## Core Features

### ✅ 100% Fluent API
Zero callbacks, pure method chaining:

```go
oas.Operation("createUser").
    RequestBody(oas.Body().ContentJSON(...)).
    Responses(
        oas.ResponseCode(201).Description("Created"),
        oas.ResponseCode(400).Description("Invalid"),
    )
```

### ✅ Type-Safe Schema Builders

```go
oas.String()    // string type
oas.Integer()   // integer type
oas.Number()    // number type
oas.Boolean()   // boolean type
oas.Array()     // array type
oas.Object()    // object type
oas.Ref(path)   // $ref to component
```

### ✅ Fluent Validation

```go
oas.String().
    MinLength(1).
    MaxLength(100).
    Pattern("^[a-zA-Z]+$").
    Format("email")

oas.Integer().
    Minimum(0).
    Maximum(100)

oas.Array().
    MinItems(1).
    UniqueItems(true).
    Items(oas.String())
```

### ✅ Clean Property Definition

```go
oas.Object().
    Required("id", "email").
    Property("id", oas.String().Format("uuid")).
    Property("email", oas.String().Format("email")).
    Property("age", oas.Integer().Minimum(18))
```

## Complete Example: E-Commerce API

```go
package main

import (
    "encoding/json"
    "os"
    "github.com/leandroluk/go/oas"
)

func main() {
    api := oas.New().
        Title("E-Commerce API").
        Description("Complete REST API for e-commerce platform").
        Version("2.0.0")

    // Product schema
    api.ComponentSchema("Product",
        oas.Object().
            Required("id", "name", "price", "stock").
            Property("id", oas.String().Format("uuid").Example("550e8400-e29b-41d4-a716-446655440000")).
            Property("name", oas.String().MinLength(3).MaxLength(100).Example("Wireless Mouse")).
            Property("description", oas.String().MaxLength(500)).
            Property("price", oas.Number().Minimum(0.01).Example(29.99)).
            Property("stock", oas.Integer().Minimum(0).Example(150)).
            Property("category", oas.Ref("#/components/schemas/Category")).
            Property("tags", oas.Array().Items(oas.String())).
            Property("active", oas.Boolean().Default(true)),
    )

    // Category schema
    api.ComponentSchema("Category",
        oas.Object().
            Required("id", "name").
            Property("id", oas.String().Example("cat_electronics")).
            Property("name", oas.String().Example("Electronics")).
            Property("parent", oas.String()),
    )

    // Product input (for creation)
    api.ComponentSchema("ProductInput",
        oas.Object().
            Required("name", "price").
            Property("name", oas.String().MinLength(3).MaxLength(100)).
            Property("description", oas.String().MaxLength(500)).
            Property("price", oas.Number().Minimum(0.01)).
            Property("stock", oas.Integer().Minimum(0).Default(0)).
            Property("categoryId", oas.String()).
            Property("tags", oas.Array().Items(oas.String())),
    )

    // Error response
    api.ComponentSchema("Error",
        oas.Object().
            Required("error", "message").
            Property("error", oas.String().Example("INVALID_INPUT")).
            Property("message", oas.String().Example("Validation failed")).
            Property("details", oas.Array().Items(
                oas.Object().
                    Property("field", oas.String()).
                    Property("issue", oas.String()),
            )),
    )

    // LIST /products
    api.Path("/products").Get(
        oas.Operation("listProducts").
            Summary("List products").
            Description("Returns paginated list of products").
            Tags("Products").
            Parameters(
                oas.InQuery("page", oas.Integer().Minimum(1).Default(1)),
                oas.InQuery("perPage", oas.Integer().Minimum(1).Maximum(100).Default(20)),
                oas.InQuery("category", oas.String()),
                oas.InQuery("search", oas.String()),
            ).
            Responses(
                oas.ResponseCode(200).
                    Description("Successful response").
                    ContentJSON(
                        oas.Array().Items(oas.Ref("#/components/schemas/Product"))
                    ),
                oas.ResponseCode(400).
                    Description("Invalid parameters").
                    ContentJSON(oas.Ref("#/components/schemas/Error")),
            ),
    )

    // CREATE /products
    api.Path("/products").Post(
        oas.Operation("createProduct").
            Summary("Create product").
            Tags("Products").
            RequestBody(
                oas.Body().
                    Required(true).
                    ContentJSON(oas.Ref("#/components/schemas/ProductInput"))
            ).
            Responses(
                oas.ResponseCode(201).
                    Description("Product created").
                    ContentJSON(oas.Ref("#/components/schemas/Product")),
                oas.ResponseCode(400).
                    Description("Validation error").
                    ContentJSON(oas.Ref("#/components/schemas/Error")),
                oas.ResponseCode(409).Description("Product already exists"),
            ),
    )

    // GET /products/{id}
    api.Path("/products/{id}").Get(
        oas.Operation("getProduct").
            Summary("Get product by ID").
            Tags("Products").
            Parameters(
                oas.InPath("id", oas.String().Format("uuid")),
            ).
            Responses(
                oas.ResponseCode(200).
                    Description("Product found").
                    ContentJSON(oas.Ref("#/components/schemas/Product")),
                oas.ResponseCode(404).
                    Description("Product not found").
                    ContentJSON(oas.Ref("#/components/schemas/Error")),
            ),
    )

    // PATCH /products/{id}
    api.Path("/products/{id}").Patch(
        oas.Operation("updateProduct").
            Summary("Update product").
            Tags("Products").
            Parameters(
                oas.InPath("id", oas.String().Format("uuid")),
            ).
            RequestBody(
                oas.Body().
                    Required(true).
                    ContentJSON(oas.Ref("#/components/schemas/ProductInput"))
            ).
            Responses(
                oas.ResponseCode(200).
                    Description("Updated successfully").
                    ContentJSON(oas.Ref("#/components/schemas/Product")),
                oas.ResponseCode(404).Description("Product not found"),
            ),
    )

    // DELETE /products/{id}
    api.Path("/products/{id}").Delete(
        oas.Operation("deleteProduct").
            Summary("Delete product").
            Tags("Products").
            Parameters(
                oas.InPath("id", oas.String().Format("uuid")),
            ).
            Responses(
                oas.ResponseCode(204).Description("Deleted successfully"),
                oas.ResponseCode(404).Description("Product not found"),
            ),
    )

    json.NewEncoder(os.Stdout).Encode(api)
}
```

## API Reference

### Schemas

#### Constructors
```go
oas.String()    // Creates string schema
oas.Integer()   // Creates integer schema
oas.Number()    // Creates number schema
oas.Boolean()   // Creates boolean schema
oas.Array()     // Creates array schema
oas.Object()    // Creates object schema
oas.Ref(path)   // Creates $ref schema
```

#### Validation Methods
```go
// String validation
.MinLength(n)
.MaxLength(n)
.Pattern(regex)
.Format("email" | "uuid" | "date" | "date-time" | ...)

// Number validation
.Minimum(n)
.Maximum(n)
.ExclusiveMinimum(bool)
.ExclusiveMaximum(bool)
.MultipleOf(n)

// Array validation
.MinItems(n)
.MaxItems(n)
.UniqueItems(bool)
.Items(schema)

// Object validation
.Required(fields...)
.Property(name, schema)
.AdditionalProperties(val)

// Common
.Example(val)
.Default(val)
.Description(text)
.ReadOnly(bool)
.WriteOnly(bool)
.Deprecated(bool)
```

#### Composition
```go
.AllOf(schemas...)  // Must match all
.OneOf(schemas...)  // Must match exactly one
.AnyOf(schemas...)  // Must match at least one
```

### Operations

```go
oas.Operation(operationId).
    Summary(text).
    Description(text).
    Tags(tags...).
    Parameters(params...).
    RequestBody(body).
    Responses(responses...)
```

### Parameters

```go
oas.InPath(name, schema)    // Path parameter (auto-required)
oas.InQuery(name, schema)   // Query parameter
oas.InHeader(name, schema)  // Header parameter
oas.InCookie(name, schema)  // Cookie parameter
```

All parameters support:
```go
.Required(bool)
.Description(text)
.Example(val)
.Deprecated(bool)
```

### Request Body

```go
oas.Body().
    Required(bool).
    Description(text).
    ContentJSON(schema)    // application/json
```

### Responses

```go
oas.ResponseCode(200)      // Exact code: "200"
oas.ResponseRange(2)       // Range: "2XX"
oas.ResponseDefault()      // Default: "default"
```

All responses support:
```go
.Description(text)         // Required
.ContentJSON(schema)       // application/json
```

### JSON Marshaling

All builders implement `json.Marshaler`:

```go
api := oas.New().Title("API").Version("1.0.0")
json.Marshal(api)  // ✅ Works directly!

schema := oas.Object().Property("id", oas.String())
json.Marshal(schema)  // ✅ Works!
```

## Best Practices

### 1. Component Reuse
Define common schemas in components:
```go
api.ComponentSchema("Error", ...)
api.ComponentSchema("PaginationMeta", ...)

// Reuse via $ref
oas.Ref("#/components/schemas/Error")
```

### 2. Input/Output Separation
```go
api.ComponentSchema("UserInput", ...)  // For POST/PATCH (without id, timestamps)
api.ComponentSchema("User", ...)       // For responses (with id, createdAt, etc.)
```

### 3. Validation Constraints
Always add reasonable constraints:
```go
oas.String().MinLength(1).MaxLength(255)  // Not unlimited
oas.Integer().Minimum(0)                  // Non-negative
oas.Array().MaxItems(100)                 // Prevent abuse
```

### 4. Use Descriptive IDs
```go
oas.Operation("createUser")    // ✅ Clear
oas.Operation("create")        // ❌ Ambiguous
```

## License

MIT
