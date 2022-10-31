package main

import (
	"syscall/js"
)

func Stub(this js.Value, args []js.Value) interface{} {
	js.Global().Call("alert", "This is not implemented yet!")

	return nil
}

func init() { // Export all methods
	RegisterFunction("RegisterToolbarMenu", RegisterToolbarMenu)
	RegisterFunction("Stub", Stub)
}
