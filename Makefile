# Makefile for ratelimiter project

# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Clean build files
clean:
	$(GOCLEAN)
	rm -f coverage.out

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# Generate mocks
mocks:
	mockery --name=RateLimiter --output=mocks

# Update dependencies
deps:
	$(GOMOD) tidy

# Lint code
lint:
	golangci-lint run

.PHONY: clean test test-coverage mocks deps lint