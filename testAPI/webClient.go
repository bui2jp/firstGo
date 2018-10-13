package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start")
	//	res, err := http.Get("https://wwww.golang.org")
	res, err := http.Get("https://www.google.co.jp")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("end")
}
