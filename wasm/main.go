package main

import (
	"fmt"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("Click", js.FuncOf(handleClick))
}

func handleClick(this js.Value, inputs []js.Value) interface{} {
	println("Button clicked")
	return nil
}

func main() {
	registerCallbacks()

	var document js.Value = js.Global().Get("document")
	var body js.Value = document.Get("body")

	var p js.Value = document.Call("createElement", "p")
	p.Set("innerHTML", "Hola")

	body.Call("appendChild", p)

	c := make(chan struct{}, 0)

	<-c

	fmt.Println("Hello world")
}
