#!/bin/bash

echo "========================================"
echo "Temperature Converter - Build Script"
echo "========================================"
echo ""

echo "[1/2] Compiling Go to WebAssembly..."
GOOS=js GOARCH=wasm go build -o main.wasm main.go

if [ $? -ne 0 ]; then
    echo ""
    echo "ERROR: Compilation failed!"
    echo "Please ensure Go is installed and in your PATH."
    echo "Download from: https://go.dev/dl/"
    exit 1
fi

echo ""
echo "[2/2] Build successful!"
echo ""
echo "========================================"
echo "Next Steps:"
echo "========================================"
echo "1. Start the server:"
echo "   python3 -m http.server 8080"
echo ""
echo "2. Open in browser:"
echo "   http://localhost:8080"
echo "========================================"
echo ""
