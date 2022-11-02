package main

import (
	"syscall/js"
)

func main() {
	RegisterToolbarMenu(js.Null(), []js.Value{js.ValueOf("Editor")})
	RegisterToolbarMenu(js.Null(), []js.Value{js.ValueOf("Plugins")})

	RegisterToolbarMenuItem(js.Null(), []js.Value{js.ValueOf("Editor"), js.ValueOf("Exclusive to Editor"), js.Global().Get("Placeholder")})
	RegisterToolbarMenuItem(js.Null(), []js.Value{js.ValueOf("Editor"), js.ValueOf("New Dialog"), js.Global().Get("ShowDialog")})
	RegisterToolbarMenuItem(js.Null(), []js.Value{js.ValueOf("Plugins"), js.ValueOf("Exclusive to Plugins"), js.Global().Get("Placeholder")})

	var tmp = make(chan int)
	<-tmp
}
