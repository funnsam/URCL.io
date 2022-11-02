package main

import (
	"strconv"
	"strings"
	"syscall/js"
)

var OpenedToolbarMenu string

type ToolbarElement struct {
	Name    string
	OnClick js.Value
}

var ToolbarElements = make(map[string][]ToolbarElement)

func OpenToobarMenu(this js.Value, args []js.Value) interface{} {
	if this.IsNull() {
		return false
	}
	if this.Get("id").String() == OpenedToolbarMenu {
		CloseToolbarMenu(js.Value{}, []js.Value{})
		return true
	}

	menu := js.Global().Get("document").Call("getElementById", "ToolbarItemSelection")
	menu.Set("innerHTML", "")
	menu.Get("style").Set("left", js.ValueOf(strconv.Itoa(this.Call("getBoundingClientRect").Get("left").Int())+"px"))

	for _, element := range ToolbarElements[this.Get("id").String()[11:]] {
		tmp := js.Global().Get("document").Call("createElement", "ul")
		tmp.Set("onclick", element.OnClick)
		tmp.Set("innerHTML", element.Name)
		menu.Call("appendChild", tmp)
	}

	js.Global().Get("document").Get("body").Call("appendChild", menu)

	OpenedToolbarMenu = this.Get("id").String()

	return true
}

func CloseToolbarMenu(this js.Value, args []js.Value) interface{} {
	OpenedToolbarMenu = ""
	menu := js.Global().Get("document").Call("getElementById", "ToolbarItemSelection")
	js.Global().Get("document").Call("getElementById", "Hidden").Call("appendChild", menu)

	return true
}

func RegisterToolbarMenu(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return false
	}

	tmp := js.Global().Get("document").Call("createElement", "toolbar-item")

	tmp.Set("innerHTML", args[0].String())
	tmp.Set("id", strings.ReplaceAll("ToolbarItem"+args[0].String(), " ", ""))
	RegisterOnClick(tmp, OpenToobarMenu)

	Toolbar.Call("appendChild", tmp)

	return true
}

func RegisterToolbarMenuItem(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return false
	}

	ToolbarElements[args[0].String()] = append(ToolbarElements[args[0].String()], ToolbarElement{args[1].String(), args[2]})
	_ = ToolbarElements

	return true
}
