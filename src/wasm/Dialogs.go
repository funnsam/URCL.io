package main

import "syscall/js"

func ShowDialog(this js.Value, args []js.Value) interface{} {
	dialog := GetElementById("Dialog")
	dialog.Get("style").Set("left", "30vw")
	dialog.Get("style").Set("top", "30vh")
	js.Global().Get("document").Get("body").Call("appendChild", dialog)
	return true
}

func HideDialog(this js.Value, args []js.Value) interface{} {
	GetElementById("Hidden").Call("appendChild", GetElementById("Dialog"))
	return true
}

func init() {
	RegisterFunction("ShowDialog", ShowDialog)
	RegisterFunction("HideDialog", HideDialog)
}
