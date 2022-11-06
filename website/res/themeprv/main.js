const URLParams = new URLSearchParams(window.location.search);

document.addEventListener('DOMContentLoaded', (event) => {
    LoadCSS("./res/spacings.css")
    LoadCSSs(URLParams.get("themelist"));
})

function LoadCSSs(list) {
    let name = list;
    let ca = name.split(',');

    for (let i = 0; i < ca.length; i++) {
        LoadCSS(ca[i])
    }
    return "";
}

function LoadCSS(url) {
    var link    = document.createElement("link");
    link.href   = url;
    link.type   = "text/css";
    link.rel    = "stylesheet";

    document.getElementsByTagName("head")[0].appendChild(link);
}