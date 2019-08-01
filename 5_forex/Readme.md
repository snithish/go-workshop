# Forex

Consume [Currency Rate](https://exchangeratesapi.io/) API.

# Folder Structure:
- `main.go`

# Building the project
```sh
make build
```
Build the micro service

# Create an executable
```sh
make run
```
Runs the service on port $PORT

# Learnings

- Consuming an API 
    - Handle 200 response from API
    - Handle 400 response from API
    - Binding response body to `struct`
- Testing
    - Custom Interface to mock struct         
- Error values
    - Differentiating failures using error type