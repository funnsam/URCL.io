function LoadWASM(url) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(url), go.importObject)
        .then((result) => go.run(result.instance));
}

LoadWASM("./res/main.wasm");
LoadWASM("./res/emulater.wasm");

if (GetCookie("theme") === "") {
    SetCookie("theme", "./themes/win7/basic.css,./themes/win7/aeropatch.css");
    location.reload();
}

LoadCSSs(GetCookie("theme"));