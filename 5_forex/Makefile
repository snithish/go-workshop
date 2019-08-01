build:
	go fmt ./...
	go vet ./...
	go build -o ./forex

run: build
	./forex

# Committing generated mocks to VCS is not advisable
generate-mocks:
	rm -rf generated_mocks
	mkdir -p generated_mocks
	mockgen -source=http_resource/http_client.go -destination=generated_mocks/mock_http_client.go -package=generated_mocks