package main

import ("fmt"
		 "bufio"
		  "os")

func main() {
	welcome := "Welcome to the program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza with ", input)
	fmt.Printf("Type of input is %T", input)

}