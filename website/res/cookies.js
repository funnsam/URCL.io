function GetCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for (let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function SetCookie(cname, cvalue) {
    const d = new Date();
    d.setMonth(d.getMonth() + 6);
    let expires = "expires="+ d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/;SameSite=Lax";
}

function ResetCookie(name) {
    SetCookie(name, "")
}

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