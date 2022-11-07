function LoadWASM(url) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(url), go.importObject)
        .then((result) => go.run(result.instance));
}

LoadWASM("./res/main.wasm");
LoadWASM("./res/emulater.wasm");

if (GetCookie("FirstTimeSetupDone") === "") {
    document.addEventListener('DOMContentLoaded', async (event) => {
    document.getElementsByTagName("body")[0].appendChild(document.getElementById("WaitForSetupBanner"))

    let SetupWizard = window.open("./firsttimesetup/", "./firsttimesetup/", "width=500,height=500");
    setInterval(async function() {if (SetupWizard.closed) {
        document.getElementById("WaitForSetupBannerText").innerHTML = "Almost there, just one more reload.";
        await new Promise(resolve => setTimeout(resolve, 500));
        location.reload();
    }}, 250);})
}

LoadCSSs(GetCookie("theme"));