# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOGENERATE := $(GOCMD) generate
BINARY_NAME := main
BUILD_DIR := build
E2E_DIR := "./tests/e2e/..."
FUNCTIONAL_DIR := "./tests/functional/..."
UNIT_DIR := "./internal/..."
CURRENCY_SERVICE := "./currency"
NOTIFIER_SERVICE := "./notifier"
GATEWAY_SERVICE := "./gateway"

all: test build

build:
	@echo "Building $(BINARY_NAME)..."
	cd $(NOTIFIER_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd
	cd $(CURRENCY_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd
	cd $(GATEWAY_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd
	# $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

unit:
	@echo "Running unit tests..."
	cd $(NOTIFIER_SERVICE) && $(GOTEST) -v $(UNIT_DIR)
	cd $(CURRENCY_SERVICE) && $(GOTEST) -v $(UNIT_DIR)

functional:
	@echo "Running functional tests..."
	cd $(NOTIFIER_SERVICE) && $(GOTEST) -v $(FUNCTIONAL_DIR)
	cd $(CURRENCY_SERVICE) && $(GOTEST) -v $(FUNCTIONAL_DIR)

e2e:
	@echo "Running end 2 end tests..."
	cd $(NOTIFIER_SERVICE) && $(GOTEST) -v $(E2E_DIR)
	cd $(CURRENCY_SERVICE) && $(GOTEST) -v $(E2E_DIR)

test:
	@echo "Running all tests..."
	$(GOTEST) $$(go list -f '{{.Dir}}/...' -m | xargs) -count=1

generate:
	@echo "Generating go code..."
	cd $(NOTIFIER_SERVICE) && $(GOGENERATE) ./... 
	cd $(CURRENCY_SERVICE) && $(GOGENERATE) ./... 

# clean:
# 	@echo "Cleaning..."
# 	cd $(NOTIFIER_SERVICE) && $(GOCLEAN) rm -f $(BUILD_DIR)/$(BINARY_NAME)
# 	cd $(CURRENCY_SERVICE) && $(GOCLEAN) rm -f $(BUILD_DIR)/$(BINARY_NAME)
# 	cd $(GATEWAY_SERVICE) && $(GOCLEAN) rm -f $(BUILD_DIR)/$(BINARY_NAME)

# run:
# 	@echo "Running $(BINARY_NAME)..."
# 	cd $(CURRENCY_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd && \
# 	$(BUILD_DIR)/$(BINARY_NAME) &
# 	cd $(NOTIFIER_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd && \
# 	$(BUILD_DIR)/$(BINARY_NAME) &
# 	cd $(GATEWAY_SERVICE) && $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd && \
# 	$(BUILD_DIR)/$(BINARY_NAME) &

# docker:
# 	@echo "Running docker..."
# 	docker build -t $(BINARY_NAME) .
# 	docker run -p 8080:8080 $(BINARY_NAME)

lint:
	@echo "Running linter..."
	golangci-lint run $$(go list -f '{{.Dir}}/...' -m | xargs)

.PHONY: all build test clean run
