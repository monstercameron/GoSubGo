#!/bin/bash

# Colors for status messages
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status messages
print_status() {
    echo -e "${YELLOW}[STATUS]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Set variables
PROJECT_ROOT=$(pwd)
WASM_OUT_DIR="$PROJECT_ROOT/static/bin"
SERVER_OUT_DIR="$PROJECT_ROOT/bin"
WASM_EXEC_JS="$PROJECT_ROOT/static/scripts/wasm_exec.js"

# Ensure output directories exist
mkdir -p "$WASM_OUT_DIR" "$SERVER_OUT_DIR" "$(dirname "$WASM_EXEC_JS")"

# Build WASM client
build_wasm() {
    print_status "Starting WASM client build..."
    
    # Check if the wasm_exec.js file exists, if not copy it from Go's source
    if [ ! -f "$WASM_EXEC_JS" ]; then
        print_status "Copying wasm_exec.js to static/scripts directory..."
        cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$WASM_EXEC_JS"
        if [ $? -eq 0 ]; then
            print_success "wasm_exec.js copied successfully to $WASM_EXEC_JS"
        else
            print_error "Failed to copy wasm_exec.js. Please copy it manually from your Go installation."
            return 1
        fi
    fi

    print_status "Compiling WASM client..."
    GOOS=js GOARCH=wasm go build -o "$WASM_OUT_DIR/main.wasm" ./cmd/client
    if [ $? -eq 0 ]; then
        print_success "WASM client built successfully."
    else
        print_error "WASM client build failed."
        return 1
    fi
}

# Build server
build_server() {
    print_status "Starting server build..."
    go build -o "$SERVER_OUT_DIR/server" ./cmd/server
    if [ $? -eq 0 ]; then
        print_success "Server built successfully."
    else
        print_error "Server build failed."
        return 1
    fi
}

# Main build process
main() {
    print_status "Beginning build process for WASM client and server..."

    # Build WASM client
    build_wasm
    if [ $? -ne 0 ]; then
        print_error "WASM client build failed. Exiting."
        exit 1
    fi

    # Build server
    build_server
    if [ $? -ne 0 ]; then
        print_error "Server build failed. Exiting."
        exit 1
    fi

    print_success "Build process completed successfully!"
    print_status "WASM client output: $WASM_OUT_DIR/main.wasm"
    print_status "wasm_exec.js location: $WASM_EXEC_JS"
    print_status "Server output: $SERVER_OUT_DIR/server"
}

# Run the main function
main