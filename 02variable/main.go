package main

import "fmt"

func main() {
	// fmt.Println("Variables in Go")
	var username string = "Vishal"
	fmt.Println("Username is", username)
	fmt.Printf("Variable type is: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable type is: %T \n", isLoggedin)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)

	var website = "www.google.com"
	fmt.Println(website)

	numberOfUsers := 100
	fmt.Println(numberOfUsers)

	
}