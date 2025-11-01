APP_NAME = cli-json-manager
CMD_PATH = ./cmd/cli-json-manager
BIN_DIR = ./bin
SRC = $(CMD_PATH)

# Default target
.PHONY: default
default: build

# Build binary into ./bin/<Application Name>
.PHONY: build
build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(SRC)

# Run using `go run` (convenient during development)
.PHONY: run
run: 
	@echo "▶️ Running from source..."
	go run $(SRC) $(ARGS)

# Run the built binary
.PHONY: run-build
run-build: build
	@echo "▶️ Running built binary..."
	$(BIN_DIR)/$(APP_NAME) $(ARGS)

# Run tests
.PHONY: test
test: 
	@echo "🧪 Running tests..."
	go test ./... -v

# Format Code (go fmt)
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	gofmt -s -w .

#  Vet /basic static checks
.PHONY: vet
vet:
	@echo "🔍 Running go vet..."
	go vet ./...

# Install binary to $GOBIN or $GOPATH/bin
.PHONY: Install
install:
	@echo "📥 Installing $(APP_NAME)..."
	go install $(SRC)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf $(BIN_DIR)

# Help
.PHONY: help
help:
	@echo "🆘 Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application from source"
	@echo "  run-build   - Run the built application"
	@echo "  test       - Run tests"
	@echo "  fmt         - Format code"
	@echo "  vet         - Run static analysis"
	@echo "  install     - Install the application"
	@echo "  clean       - Clean build artifacts"
