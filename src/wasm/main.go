package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	RegisterToolbarMenu(js.Null(), []js.Value{js.ValueOf("Editor")})

	var tmp = make(chan int)
	<-tmp
}

func OpenToobarMenu(this js.Value, args []js.Value) interface{} {
	if this.IsNull() {
		return false
	}

	menu := js.Global().Get("document").Call("getElementById", "ToolbarItemSelection")

	menu.Get("style").Set("left", js.ValueOf(strconv.Itoa(this.Call("getBoundingClientRect").Get("left").Int())+"px"))

	js.Global().Get("document").Get("body").Call("appendChild", menu)

	return true
}

func CloseToolbarMenu(this js.Value, args []js.Value) interface{} {
	menu := js.Global().Get("document").Call("getElementById", "ToolbarItemSelection")
	js.Global().Get("document").Call("getElementById", "Hidden").Call("appendChild", menu)

	return true
}

func init() {
	RegisterFunction("OpenToobarMenu", OpenToobarMenu)

	RegisterOnClick("CloseToolbarItemSelection", CloseToolbarMenu)
}
