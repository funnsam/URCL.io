package main

import "syscall/js"

var Toolbar = js.Global().Get("document").Call("getElementById", "Toolbar")
