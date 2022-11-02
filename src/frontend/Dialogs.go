package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
)

func SimpleShowDialog(this js.Value, args []js.Value) interface{} {
	SetDialogContent(this, []js.Value{js.ValueOf("simple"), args[0], args[1]})
	ShowDialog(this, args)
	return true
}

func SetDialogContent(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return false
	}

	switch args[0].String() {
	case "simple":
		GetElementById("DialogTitle").Set("innerHTML", args[1].String())
		GetElementById("DialogContent").Set("innerHTML", args[2].String())

	default:
		return false
	}

	return true
}

func ShowDialog(this js.Value, args []js.Value) interface{} {
	dialog := GetElementById("Dialog").Call("cloneNode", true)
	dialogid := strconv.Itoa(rand.Int())

	dialog.Set("id", "Dialog"+dialogid)
	dialog.Get("childNodes").Index(3).Set("id", "DialogTitle"+dialogid)
	dialog.Get("childNodes").Index(6).Set("id", "DialogContent"+dialogid)
	dialog.Get("style").Set("left", "30vw")
	dialog.Get("style").Set("top", "30vh")
	js.Global().Get("document").Get("body").Call("appendChild", dialog)
	return true
}

func CloseDialog(this js.Value, args []js.Value) interface{} {
	args[0].Call("remove")
	return true
}
