.PHONY: help build test test-coverage lint fmt vet clean install-tools

# Default target
help:
	@echo "Available targets:"
	@echo "  build          - Build the project"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  lint           - Run golangci-lint"
	@echo "  fmt            - Format code with gofmt"
	@echo "  vet            - Run go vet"
	@echo "  clean          - Remove build artifacts"
	@echo "  install-tools  - Install development tools"

# Build the project
build:
	go build -v ./...

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run golangci-lint (requires golangci-lint to be installed)
lint:
	golangci-lint run ./...

# Format code
fmt:
	gofmt -s -w .

# Run go vet
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -f coverage.txt coverage.html
	rm -rf dist/ build/ bin/
	go clean

# Install development tools
install-tools:
	@echo "Installing golangci-lint..."
	@which golangci-lint > /dev/null || \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development tools installed"

# Run all checks before commit
pre-commit: fmt vet lint test
	@echo "All checks passed!"
