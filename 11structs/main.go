package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")

	vishal := User{"Vishal", "vishal@go.dev", 21, true}
	fmt.Println(vishal)
	fmt.Printf("Vishal details are: %+v\n", vishal)
}

type User struct {
	Name string
	Email string
    Age int
	Status bool
}

