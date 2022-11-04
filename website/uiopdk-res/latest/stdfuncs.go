package main

import "syscall/js"

func main() {
	var tmp = make(chan int)
	<-tmp
}

func RegisterFunction(name string, function func(this js.Value, args []js.Value) interface{}) {
	js.Global().Set(name, js.FuncOf(function))
}

func RegisterOnClick(tmp js.Value, function func(this js.Value, args []js.Value) interface{}) {
	tmp.Set("onclick", js.FuncOf(function))
}

// Same as getElementById in JavaScript
func GetElementByID(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func RegisterToolbarMenu(name string) {
	js.Global().Get("RegisterToolbarMenu").Invoke(name)
}

func RegisterToolbarMenuItem(menu, name string, jsfunc js.Func) {
	js.Global().Get("RegisterToolbarMenuItem").Invoke(menu, name, jsfunc)
}

func ShowDialog(tod, title, content string) int {
	return js.Global().Get("ShowDialog").Invoke(tod, title, content).Int()
}
