package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	mynumber := 25

	var ptr = &mynumber
	fmt.Println("The value of ptr is", ptr)
	fmt.Println("The value of ptr is", *ptr)
	*ptr = *ptr + 2
	fmt.Println("The value is", mynumber)
}
