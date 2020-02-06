package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe("0.0.0.0:8800", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello world")
	fmt.Fprintf(writer, "HELLO WORLD")
}
