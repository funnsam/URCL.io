const URLParams = new URLSearchParams(window.location.search);

document.addEventListener('DOMContentLoaded', (event) => {
    document.getElementById("tmp").innerHTML = `"` + URLParams.get("themelist") + `"`;
})