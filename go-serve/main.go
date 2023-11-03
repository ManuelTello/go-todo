package main

import (
	"fmt"
	"net/http"
)

func main(){
    fs := http.FileServer(http.Dir("../go-wasm/"))

    http.Handle("/",fs)

    fmt.Println("Serving at port :3030")
    http.ListenAndServe(":3030",nil)
}
