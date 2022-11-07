import {OnStartThemeSetup} from "./themesetup.js"

export async function FirstTimeSetupSetup() {
    await OnStartThemeSetup();
    await SetCookie("FirstTimeSetupDone", "yes");
}