# go-query-builder
SQL Query Builder written in Go for sql module

[![GitHub issues](https://img.shields.io/github/issues/c0de4un/go-query-builder)](https://github.com/c0de4un/go-query-builder/issues)
[![GitHub stars](https://img.shields.io/github/stars/c0de4un/go-query-builder)](https://github.com/c0de4un/go-query-builder/stargazers)
[![GitHub license](https://img.shields.io/github/license/c0de4un/go-query-builder)](https://github.com/c0de4un/go-query-builder/blob/main/LICENSE)
![GitHub issues](https://img.shields.io/badge/language-Go-blue)

# Requirements
* Go 1.17+

# Tests
    ```go
    $go test
    ```
# Examples

Simple select
```go
package main

import (
    import "github.com/c0de4un/go-query-builder"
    // ...
)

func main() {
	builder := NewBuilder()
	builder.Select("users.*")
	builder.Select("projects.id")
    builder.From("users")
    builder.Join("projects", "projects.id = users.project_id", "LEFT")
    builder.Where("users.id", ">", 1)
    query := builder.Compile()

    // Execute query
}
```