package main

import (
	"syscall/js"
)

func Placeholder(this js.Value, args []js.Value) interface{} {
	SimpleShowDialog(js.Null(), []js.Value{js.ValueOf("Fallback Dialog"), js.ValueOf("Is it not implemented? idk.")})

	return nil
}

func PlaceholderDog(this js.Value, args []js.Value) interface{} {
	SimpleShowDialog(js.Null(), []js.Value{js.ValueOf("Dauog"), js.ValueOf(`<img src="http:/` + `/placecorgi.com/256">`)})

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
	RegisterFunction("PlaceholderDog", PlaceholderDog)
}
