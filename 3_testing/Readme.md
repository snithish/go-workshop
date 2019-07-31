# Unit Tests in Go

In this module we will need to implement string reverser.

# Folder Structure:
- `main.go` invokes string reverser with arbitrary input
    - pay attention to the import path for the string reverser
- the reverser is present in `utils` package
- implementation goes into `reverser.go` 
- tests are written in `reverser_table_test.go` and `reverser_test.go`

# Running Tests
```sh
go test ./...
```

Runs test recursively in all sub packages

# Learnings
- test file naming convention
- table driven tests
- packages
- top notch unicode support in Go
    - runes
    - bytes
