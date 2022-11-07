function LoadWASM(url) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(url), go.importObject)
        .then((result) => go.run(result.instance));
}

LoadWASM("./res/main.wasm");
LoadWASM("./res/emulater.wasm");

LoadCSSs(GetCookie("theme"));

import {FirstTimeSetupSetup} from "./firsttimesetup/setup.js"; 

if (GetCookie("FirstTimeSetupDone") === "") {
    document.addEventListener('DOMContentLoaded', async (event) => {
        document.getElementsByTagName("body")[0].appendChild(document.getElementById("WaitForSetupBanner"))
        await FirstTimeSetupSetup();
        location.reload();
    })
}