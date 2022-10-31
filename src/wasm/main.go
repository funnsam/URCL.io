package main

import (
	"syscall/js"
)

func main() {
	RegisterToolbarMenu(js.Null(), []js.Value{js.ValueOf("Editor")})
	RegisterToolbarMenu(js.Null(), []js.Value{js.ValueOf("Plugins")})

	RegisterToolbarMenuItem(js.Null(), []js.Value{js.ValueOf("Editor"), js.ValueOf("Exclusive hello world"), js.Global().Get("Stub")})

	var tmp = make(chan int)
	<-tmp
}
