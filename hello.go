package main

import (
	"fmt"
)

func main() {

	fmt.Println("Go言語はじめました！start")

	var myCount int
	myCount = 100

	for i := 0; i < myCount; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	var msg2 = "123"

	msg2 = "hello"

	fmt.Println(msg2)

	var msg string

	msg = "test test"

	fmt.Println(msg)

	go myFunc01()
	go myFunc01()

	fmt.Println("Go言語はじめました！end")

}
func myFunc01() {
	//time.Sleep(3 * time.Second)
	fmt.Println("my func")
}
