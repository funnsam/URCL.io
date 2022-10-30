package main

import (
	"strings"
	"syscall/js"
)

func RegisterToolbarMenu(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return false
	}
	tmp := js.Global().Get("document").Call("createElement", "toolbar-item")

	tmp.Set("innerHTML", args[0].String())
	tmp.Set("id", strings.ReplaceAll("ToolbarItem"+args[0].String(), " ", ""))
	tmp.Set("onclick", js.ValueOf(js.FuncOf(OpenToobarMenu)))

	Toolbar.Call("appendChild", tmp)

	return true
}

func Stub(this js.Value, args []js.Value) interface{} {
	js.Global().Call("alert", "This is not implemented yet!")

	return nil
}

func init() { // Export all methods
	RegisterFunction("RegisterToolbarMenu", RegisterToolbarMenu)
	RegisterFunction("Stub", Stub)
}
