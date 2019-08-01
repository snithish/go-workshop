# Forex

1. Consume [Currency Rate](https://exchangeratesapi.io/) API.
2. Handle 400 from API appropriately
3. Complete tests for HTTP Resource using mocks  


# Project Structure
- Package `http_resource_hard_test` contains a HTTP Resource implementation which is hard to test
- Package `http_resource` contains a HTTP Resource implementation which can be easily tested with Mocking

# Binary Dependency for generating mocks
```sh
go install github.com/golang/mock/mockgen
```

# Building the project
```sh
make build
```
Build the micro service

# Running the micro service
```sh
make run
```
Runs the service on port 8080

# Learnings

- Consuming an API 
    - Handle 200 response from API
    - Handle 400 response from API
    - Unmarshalling response body
- Testing
    - Custom Interface to mock struct
    - If we could not do this, then we end up having to stub the calls 
- Error values
    - Differentiating failures using error type
    
# Exercise
- Binding validators
    - `AmountToConvert` must be greater than 0
- Write tests using httptest
    - Non trivial