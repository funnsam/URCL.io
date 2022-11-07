document.addEventListener('DOMContentLoaded', async (event) => {
    await OnStartThemeSetup();
    SetCookie("FirstTimeSetupDone", "yes");
    window.close();
})

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
    let expires = "expires="+ new Date(32530498344809).toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/;SameSite=Lax";
}

function ResetCookie(name) {
    SetCookie(name, "")
}