package main

import (
	"syscall/js"
)

func RegisterFunction(name string, function func(this js.Value, args []js.Value) interface{}) {
	js.Global().Set(name, js.FuncOf(function))
}

func RegisterOnClick(id string, function func(this js.Value, args []js.Value) interface{}) {
	js.Global().Get("document").Call("getElementById", id).Set("onclick", js.FuncOf(function))
}
