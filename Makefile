# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
BINARY_NAME := exchange
BUILD_DIR := build

all: test build

build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/web

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run:
	@echo "Running $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/web
	$(BUILD_DIR)/$(BINARY_NAME)

docker:
	@echo "Running docker..."
	docker build -t $(BINARY_NAME) .
	docker run -p 8080:8080 $(BINARY_NAME)

.PHONY: all build test clean run
