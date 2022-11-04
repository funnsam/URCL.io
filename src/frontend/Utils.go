package main

import (
	"syscall/js"
)

func RegisterFunction(name string, function func(this js.Value, args []js.Value) interface{}) {
	js.Global().Set(name, js.FuncOf(function))
}

func RegisterOnClick(tmp js.Value, function func(this js.Value, args []js.Value) interface{}) {
	tmp.Set("onclick", js.FuncOf(function))
}

func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func Hang() {
	var tmp = make(chan int)
	<-tmp
}
