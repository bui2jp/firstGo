package main

import (
	"fmt"
	"net/http"
)

var myApis []string

func main() {
	fmt.Println("start")
	myApis = make([]string, 0)
	http.HandleFunc("/myApi", myApiFunc)
	http.ListenAndServe(":8080", nil)
	fmt.Println("end")
}

func myApiFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "string from myApi Func")
}
