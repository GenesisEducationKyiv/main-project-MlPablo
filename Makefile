# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOGENERATE := $(GOCMD) generate
BINARY_NAME := exchange
BUILD_DIR := build
E2E_DIR := "./tests/e2e/..."
FUNCTIONAL_DIR := "./tests/functional/..."
UNIT_DIR := "./internal/..."

all: test build

build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

unit:
	@echo "Running unit tests..."
	$(GOTEST) -v $(UNIT_DIR)

functional:
	@echo "Running functional tests..."
	$(GOTEST) -v $(FUNCTIONAL_DIR)

e2e:
	@echo "Running end 2 end tests..."
	$(GOTEST) -v $(E2E_DIR)

test:
	@echo "Running all tests..."
	$(GOTEST) $$(go list -f '{{.Dir}}/...' -m | xargs) -count=1

generate:
	@echo "Generating go code..."
	$(GOGENERATE) ./...


clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run:
	@echo "Running $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd
	$(BUILD_DIR)/$(BINARY_NAME)

docker:
	@echo "Running docker..."
	docker build -t $(BINARY_NAME) .
	docker run -p 8080:8080 $(BINARY_NAME)

lint:
	@echo "Running linter..."
	golangci-lint run $$(go list -f '{{.Dir}}/...' -m | xargs)

.PHONY: all build test clean run
