package main

import (
	"syscall/js"
)

func Placeholder(this js.Value, args []js.Value) interface{} {
	ShowDialog(js.Null(), []js.Value{})

	return nil
}

func init() { // Export all methods
	RegisterFunction("RegisterToolbarMenu", RegisterToolbarMenu)
	RegisterFunction("OpenToobarMenu", OpenToobarMenu)
	RegisterFunction("RegisterToolbarMenuItem", RegisterToolbarMenuItem)

	RegisterFunction("SimpleShowDialog", SimpleShowDialog)
	RegisterFunction("SetDialogContent", SetDialogContent)
	RegisterFunction("ShowDialog", ShowDialog)
	RegisterFunction("CloseDialog", CloseDialog)

	RegisterFunction("Placeholder", Placeholder)
}
