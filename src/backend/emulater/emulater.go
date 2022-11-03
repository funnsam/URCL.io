package main

import "syscall/js"

func InvokeEmulater(this js.Value, args []js.Value) interface{} {
	Tokenize([]byte(GetElementById("Codes").Get("value").String()))
	return true
}

func init() {
	RegisterFunction("InvokeEmulater", InvokeEmulater)
}
