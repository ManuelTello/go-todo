package main

import "net/http"

func main(){
    fs := http.FileServer(http.Dir("../client/"))

    http.Handle("/",fs)

    http.ListenAndServe(":3030",nil)
}
