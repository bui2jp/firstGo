package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("end by defer")
	fmt.Println("start")
	selectUser()
	fmt.Println("end")
}
