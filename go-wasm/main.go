//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func registerCallbacks(){
    js.Global().Set("Click",js.FuncOf(handleClick))
}

func handleClick(this js.Value,inputs []js.Value)interface{}{
    println("Button clicked")
    return nil
}

func main(){
    fmt.Println("Hello world")
}
