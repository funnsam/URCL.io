package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
)

// A simple wrapper function for ShowDialog()
//
// Deprecated: Will be featured in the library
func SimpleShowDialog(this js.Value, args []js.Value) interface{} {
	return ShowDialog(this, []js.Value{js.ValueOf("simple"), args[0], args[1]})
}

// Usage: nil.SetDialogContent(dialogID, contentType, title, content)
func SetDialogContent(this js.Value, args []js.Value) interface{} {
	if len(args) < 4 {
		return false
	}

	switch args[1].String() {
	case "simple":
		GetElementById("DialogTitle"+strconv.Itoa(args[0].Int())).Set("innerHTML", args[2].String())
		GetElementById("DialogContent"+strconv.Itoa(args[0].Int())).Set("innerHTML", args[3].String())

	default:
		return false
	}

	return true
}

// Usage: nil.SetDialogContent(contentType, title, content)
func ShowDialog(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return 0
	}

	dialog := GetElementById("Dialog").Call("cloneNode", true)
	dialogidint := rand.Int()
	dialogid := strconv.Itoa(dialogidint)

	dialog.Set("id", "Dialog"+dialogid)
	dialog.Get("childNodes").Index(3).Set("id", "DialogTitle"+dialogid)
	dialog.Get("childNodes").Index(6).Set("id", "DialogContent"+dialogid)
	dialog.Get("style").Set("left", "30vw")
	dialog.Get("style").Set("top", "30vh")

	switch args[0].String() {
	case "simple":
		dialog.Get("childNodes").Index(3).Set("innerHTML", args[1].String())
		dialog.Get("childNodes").Index(6).Set("innerHTML", args[2].String())

	default:
		return 0
	}

	js.Global().Get("document").Get("body").Call("appendChild", dialog)
	return dialogidint
}

// Usage: dialog.CloseDialog()
func CloseDialog(this js.Value, args []js.Value) interface{} {
	args[0].Call("remove")
	return true
}
