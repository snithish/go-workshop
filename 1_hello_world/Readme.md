# Genesis

Starting point of any language learning.

# Folder Structure:
- `main.go`

# Executing Go File
```sh
go run main.go
```
Executes a Go file

# Create an executable
```sh
go build
./hello_world
```
Build (create an executable) and then run it

# Learnings
- `main function` 
    - entry point for programming execution  
    - should be in main package
- import external package `fmt` for printing to console

---

**Exercise**
- What happens when we try to run compiled binary in another OS?
- How to build OS specific binaries?

```sh
GOOS=linux go build
./hello_world
```
**_exec format error: ./hello_world_**