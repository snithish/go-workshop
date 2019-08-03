# Pet Store

An example documenting folder structure and code organization for a Gin webservice.
    
# Binary Dependencies

```sh
brew install richgo
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.17.1
```

# Build Project
```sh
make generate-mocks build
```

# Start service
```sh
make generate-mocks run
```

# Learnings
- define routes and handlers -> Controllers 
- Reading reading request body, path param and query param
- Request body JSON to struct using `json tags`
- Model validations
- Mocking dependencies 
- import package `fmt` for printing to console

---

**Exercise**
- Fix Linter errors `curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.17.1`
- Create tests for pet_service
- Create tests for pet_repository
- Instead of `Validate()` method on models use https://github.com/gin-gonic/gin#custom-validators