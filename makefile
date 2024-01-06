# Go variables
GO := go
BUILD := $(GO) build

# Application details
APP_NAME := asm-fmt
OUTPUT_DIR := build
SRC_FILE := cmd/main.go

# Cross-compilation variables
ARCH := $(shell go env GOARCH)
OS := $(shell go env GOOS)

# Main build target
build: clean
	@echo "Building for $(OS) $(ARCH)"
	@$(BUILD) -o $(OUTPUT_DIR)/$(APP_NAME) $(SRC_FILE)

# Clean target
clean:
	@echo "Cleaning..."
	# @rm -rf $(OUTPUT_DIR)

# Other targets
.PHONY: clean