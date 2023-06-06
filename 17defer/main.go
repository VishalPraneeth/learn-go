package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("defer")
	defer fmt.Println("defer2") // defer will be executed in LIFO order
	fmt.Println("end")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}