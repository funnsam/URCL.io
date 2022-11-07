export async function OnStartThemeSetup() {
    const SystemSpeed = await BenchmarkSystemAcc();
    
    if (SystemSpeed >= 4444.08) {
        SetCookie("theme", "./themes/win7/basic.css")
    } else {
        SetCookie("theme", "./themes/win7/basic.css,./themes/win7/aeropatch.css")
    }
}

function Delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}

async function BenchmarkSystemAcc() {
    const First = BenchmarkSystem();
    await Delay(3000);

    const Second = BenchmarkSystem();
    await Delay(3000);

    const Third = BenchmarkSystem();

    return (First + Second + Third) / 3;
}

function BenchmarkSystem() {
    const BenchmarkLevel = 2500000000;

    const Before = performance.now();
    for (let i = BenchmarkLevel; i > 0; i--) {}
    const After = performance.now();

    return After - Before   // 7777.14146 for 1GHz
}