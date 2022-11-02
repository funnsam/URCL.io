function LoadWASM(url) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(url), go.importObject)
        .then((result) => go.run(result.instance));
}

LoadWASM("./res/main.wasm")