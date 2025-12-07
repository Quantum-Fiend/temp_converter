@echo off
echo ========================================
echo Temperature Converter - Build Script
echo ========================================
echo.

echo [1/2] Compiling Go to WebAssembly...
set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go

if %ERRORLEVEL% NEQ 0 (
    echo.
    echo ERROR: Compilation failed!
    echo Please ensure Go is installed and in your PATH.
    echo Download from: https://go.dev/dl/
    pause
    exit /b 1
)

echo.
echo [2/2] Build successful!
echo.
echo ========================================
echo Next Steps:
echo ========================================
echo 1. Start the server:
echo    python -m http.server 8080
echo.
echo 2. Open in browser:
echo    http://localhost:8080
echo ========================================
echo.
pause
