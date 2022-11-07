let ThemeList = GetCookie("theme")
let ThemeSelMode = 0 // 0=Normal, 1=Advanced

function SetupSettingMenu() {
    RegisterToolbarMenuItem("Editor", "Settings", function() {
        ShowDialog("simple", "Settings", `Theme:<br>
        <button onclick="ShowThemingWindow();">Select a theme here</button>
        <br>
        <button onclick="location.reload()" class="DialogFooter">Reload Page</button>`);
    });
}

function ShowThemingWindow() {
    ShowDialog("simplewithsize", "Select a theme", `
        <iframe src="./themepreview.html?themelist=${ThemeList}" style="width: calc(100% - 5px); height: 450px" id="ThemePreviewWindow"></iframe><br>

        <select id="ThemeSel" onchange="UserFriendlyChangeTheme();">
        <option>--Chose a theme--</option>
        <option>Windows 7 Basic</option>
        <option>Windows 7 Aero</option>
        <option>Dreaming</option>
        <option>Modern UI</option>
        </select><br>
        <lable>Set your own: </lable><input onkeydown="AdavncedSetTheme(event, this.value);" id="AdvancedTheme"
        style="width: calc(100% - 8px);"></input><br><button onclick="ChangeTheme();CloseDialog(this.parentNode.parentNode);" class="DialogFooter">
        Apply theme (You need to restart to apply)</button>`, "1000px", "750px")
}

function GetThemeSources(name) {
    switch (name) {
        case "Windows 7 Basic":
            return "./themes/win7/basic.css";
        case "Windows 7 Aero":
            return "./themes/win7/basic.css,./themes/win7/aeropatch.css";
        case "Dreaming":
            return "./themes/dream/basic.css";
        case "Modern UI":
            return "./themes/modern/modern.css";
    }
}

function UserFriendlyChangeTheme() {
    ThemeList = GetThemeSources(document.getElementById("ThemeSel").options[document.getElementById("ThemeSel").selectedIndex].text);
    document.getElementById("ThemePreviewWindow").src = `./themepreview.html?themelist=${ThemeList}`;

    ThemeSelMode = 0;
}

let ReadyToRestartSetup = false

function AdavncedSetTheme(ev, value) {
    switch (ev.keyCode) {
        case 17:
            ReadyToRestartSetup = true;
            break;

        case 68:
            if (ReadyToRestartSetup == true) ResetCookie("FirstTimeSetupDone");
            break;

        default:
            document.getElementById("ThemePreviewWindow").src = `./themepreview.html?themelist=${value}`;
            document.getElementById("ThemeSel").selectedIndex = 0;

            ReadyToRestartSetup = false;
            ThemeSelMode = 1;
            break;
    }
}

function ChangeTheme() {
    switch (ThemeSelMode) {
        case 0:
            const name = document.getElementById("ThemeSel").options[document.getElementById("ThemeSel").selectedIndex].text;
            SetCookie("theme", GetThemeSources(name));
            break;
            
        case 1:
            SetCookie("theme", document.getElementById("AdvancedTheme").value);
            break;
    }
}