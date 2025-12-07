// Minimal Wasm loader - all logic is in Go
const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
  })
  .catch((err) => {
    console.error("Failed to load Wasm:", err);
    document.getElementById('result').innerText = "⚠️ Failed to load. Please run: python -m http.server";
    document.getElementById('result').classList.add('error');
  });
