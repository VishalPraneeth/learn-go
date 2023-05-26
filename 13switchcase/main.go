package main

import ("fmt"
		"math/rand"
		"time")

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)
	switch  diceNumber {
		case 1: 
		
            fmt.Println("You rolled 1")
		case 2:
			fmt.Println("You rolled 2")
		
        case 3:
			fmt.Println("You rolled 3")
		
        case 4:
			fmt.Println("You rolled 4")
		
        case 5:
			fmt.Println("You rolled 5")
		
		case 6:
			fmt.Println("You rolled 6")

		default:
			fmt.Println("You rolled unknown number")	
		
	}
	 
}