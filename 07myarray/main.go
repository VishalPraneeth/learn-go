package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GoLang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	//fruitList[2] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List Length:", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Veg List:", len(vegList))
}