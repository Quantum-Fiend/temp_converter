# Temperature Converter

A modern, dark-themed temperature converter built with **Go + WebAssembly** featuring stunning animations and glassmorphism design.

![Dark Theme](https://img.shields.io/badge/theme-dark-000000)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?logo=webassembly&logoColor=white)

## ✨ Features

- 🔥 **100% Go Logic** - All conversion logic and DOM manipulation in Go
- 🌙 **Premium Dark UI** - Glassmorphism with animated gradients
- ✨ **Smooth Animations** - Floating particles, glowing effects, transitions
- ⚡ **WebAssembly** - Fast, efficient client-side execution
- 📱 **Responsive** - Works on all screen sizes
- ⌨️ **Keyboard Support** - Press Enter to convert

## 🚀 Quick Start

### Prerequisites

- [Go](https://go.dev/dl/) (1.16 or higher)
- Python (for local server)

### Build & Run

**Windows:**
```cmd
build.bat
```

**Linux/Mac:**
```bash
chmod +x build.sh
./build.sh
```

**Manual Build:**
```bash
# Windows PowerShell
$env:GOOS='js'; $env:GOARCH='wasm'; go build -o main.wasm main.go

# Linux/Mac
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

**Start Server:**
```bash
python -m http.server 8080
```

**Open Browser:**
```
http://localhost:8080
```

## 📁 Project Structure

```
temp_converter/
├── main.go          # Go logic + DOM handling
├── index.html       # HTML structure
├── styles.css       # Dark animated styles
├── index.js         # Wasm loader (minimal)
├── wasm_exec.js     # Go Wasm runtime
├── build.bat        # Windows build script
├── build.sh         # Unix build script
└── README.md        # This file
```

## 🎨 Design Highlights

- **Animated Background** - Shifting gradient with floating particles
- **Glassmorphism** - Frosted glass effect with backdrop blur
- **Neon Accents** - Cyan (#00f5ff) and purple (#7b2ff7)
- **Smooth Transitions** - All interactions are animated
- **Glow Effects** - Text shadows and button glows
- **Responsive Layout** - Mobile-friendly design

## 🛠️ Technology Stack

- **Backend:** Go (WebAssembly)
- **Frontend:** HTML5, CSS3, Vanilla JS
- **Build:** Go compiler
- **Server:** Python HTTP server (development)

## 📝 How It Works

1. **Go Code** (`main.go`) compiles to WebAssembly
2. **Wasm Runtime** (`wasm_exec.js`) loads the binary
3. **Go Functions** handle all DOM events and logic
4. **CSS Animations** provide visual effects
5. **Minimal JS** only loads Wasm, everything else is Go

## 🔧 Development

The entire application logic is in Go:
- DOM manipulation via `syscall/js`
- Event listeners (click, keypress)
- Input validation
- Temperature conversion
- Error handling

## 📄 License

MIT License - feel free to use this project however you'd like!

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit PRs.

---

**Made with ❤️ using Go + WebAssembly**
