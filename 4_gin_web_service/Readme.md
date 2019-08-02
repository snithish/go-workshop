# Pet Store

An example documenting folder structure and code organization for a Gin webservice.
    
# Build Project
```sh
make build
```

# Start service
```sh
make run
```

# Learnings
- define routes and handlers -> Controllers 
- Reading reading request body, path param and query param
- Request body JSON to struct using `json tags`
- Model validations
- Mocking dependencies 
- import external package `fmt` for printing to console

---

**Exercise**
- Instead of `Validate()` method on models use https://github.com/gin-gonic/gin#custom-validators