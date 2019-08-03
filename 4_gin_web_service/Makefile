build:
	golangci-lint run
	go fmt ./...
	go vet ./...
	richgo test ./... -v
	go build -o ./bin/pet_store

run: build
	./bin/pet_store

# Committing generated mocks to VCS is not advisable
generate-mocks:
	rm -rf generated_mocks
	mkdir -p generated_mocks
	mockgen -source=services/pet_service.go -destination=generated_mocks/mock_pet_service.go -package=generated_mocks
	mockgen -source=repositories/pet_repository.go -destination=generated_mocks/mock_pet_repository.go -package=generated_mocks