var ThemeList = GetCookie("theme")

function SetupSettingMenu() {
    RegisterToolbarMenuItem("Editor", "Settings", function() {
        ShowDialog("simple", "Settings", `Theme:<br>
        <button onclick="ShowThemingWindow();">Select a theme here</button>
        <br>
        <button onclick="location.reload()" class="DialogFooter">Reload Page</button>`);
    });
}

function ShowThemingWindow() {
    ShowDialog("simple", "Select a theme", `
        <iframe src="./themepreview.html?themelist=${ThemeList}" style="width: calc(100% - 5px); height: 250px" id="ThemePreviewWindow"></iframe><br>

        <select id="ThemeSel" onchange="SetThemePreview();">
        <option>--Chose a theme--</option>
        <option>Windows 7 Basic</option>
        <option>Windows 7 Aero</option>
        <option>Modern UI</option>
        </select><br>
        <button onclick="ChangeTheme();CloseDialog(this.parentNode.parentNode);" class="DialogFooter">Apply theme (You need to restart to apply)</button>`)
}

function GetThemeSources(name) {
    switch (name) {
        case "Windows 7 Basic":
            return "./themes/win7/basic.css";
        case "Windows 7 Aero":
            return "./themes/win7/basic.css,./themes/win7/aeropatch.css";
        case "Modern UI":
            return "./themes/modern/modern.css";
    }
}

function SetThemePreview() {
    ThemeList = GetThemeSources(document.getElementById("ThemeSel").options[document.getElementById("ThemeSel").selectedIndex].text);
    document.getElementById("ThemePreviewWindow").src = `./themepreview.html?themelist=${ThemeList}`;
}

function ChangeTheme() {
    const name = document.getElementById("ThemeSel").options[document.getElementById("ThemeSel").selectedIndex].text;
    SetCookie("theme", GetThemeSources(name));
}