const go = new Go();
WebAssembly.instantiateStreaming(fetch("./res/main.wasm"), go.importObject)
    .then((result) => go.run(result.instance));