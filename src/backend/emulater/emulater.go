package main

import "syscall/js"

func InvokeEmulater(this js.Value, args []js.Value) interface{} {
	return true
}

func main() {
	Hang()
}

func init() {
	RegisterFunction("InvokeEmulater", InvokeEmulater)
}
