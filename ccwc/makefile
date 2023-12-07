BINARY_NAME=ccwc
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin
GOFLAGS=-ldflags="-s -w"

default: build install

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@install -m 755 $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)..."
	@rm -f $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: build install uninstall clean test
