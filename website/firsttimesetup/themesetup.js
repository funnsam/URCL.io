async function OnStartThemeSetup() {
    const Bar = document.getElementById("ThemeSetupBar");
    const SystemSpeed = await BenchmarkSystemAcc();
    
    if (SystemSpeed >= 4444.08) {
        SetCookie("theme", "./themes/win7/basic.css")
    } else {
        SetCookie("theme", "./themes/win7/basic.css,./themes/win7/aeropatch.css")
    }

    Bar.value = 4;
}

function Delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}

async function BenchmarkSystemAcc() {
    const Bar = document.getElementById("ThemeSetupBar");

    const First = BenchmarkSystem();
    await Delay(3000);
    Bar.value = 1;

    const Second = BenchmarkSystem();
    await Delay(3000);
    Bar.value = 2;

    const Third = BenchmarkSystem();
    Bar.value = 3;

    return (First + Second + Third) / 3;
}

function BenchmarkSystem() {
    const BenchmarkLevel = 2500000000;

    const Before = performance.now();
    for (let i = BenchmarkLevel; i > 0; i--) {}
    const After = performance.now();

    return After - Before   // 7777.14146 for 1GHz
}