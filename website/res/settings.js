function SetupSettingMenu() {
    RegisterToolbarMenuItem("Editor", "Settings", function() {
        ShowDialog("simple", "Settings", `Theme:<br><select onchange="ChangeTheme(this)">
        <option>--Chose a theme--</option>
        <option>Windows 7 Basic</option>
        <option>Windows 7 Aero</option>
        <option>Modern UI</option>
        </select><br>
        <button onclick="location.reload()">Reload Page</button>`);
    });
}

function ChangeTheme(that) {
    switch (that.options[that.selectedIndex].text) {
        case "Windows 7 Basic":
            SetCookie("theme", "./themes/win7/basic.css,");
            break;
        case "Windows 7 Aero":
            SetCookie("theme", "./themes/win7/basic.css,./themes/win7/aeropatch.css");
            break;
        case "Modern UI":
            SetCookie("theme", "./themes/modern/modern.css");
            break;
    }
}