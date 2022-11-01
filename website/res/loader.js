const go = new Go();
WebAssembly.instantiateStreaming(fetch("./res/main.wasm"), go.importObject)
    .then((result) => go.run(result.instance));

function StartDrag(ev) {
    const rect = ev.target.getBoundingClientRect();
  
    offsetX = ev.clientX - rect.x;
    offsetY = ev.clientY - rect.y;
};

function StopDrag(ev) {
    document.getElementById("Dialog").style.left = ev.clientX - offsetX + 'px';
    document.getElementById("Dialog").style.top = ev.clientY - offsetY + 'px';
};