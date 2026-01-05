# Package Set

A utility package for Go that provides a generic `Field[T]` type to handle **Partial Updates** and **Patch** operations via JSON.

## Features

- **Generic `Field[T]`**: Works with any data type.
- **Explicit Tracking**: Knows exactly if a field was present in the JSON payload.
- **SQL Ready**: Convert structs to maps for `UPDATE` operations using `ToMap`.
- **JSON Tag Support**: Respects standard `json` tags for key naming.

## Usage

### Define your DTO
```go
type UserUpdate struct {
    Name  set.Field[string] `json:"name"`
    Age   set.Field[int]    `json:"age"`
}