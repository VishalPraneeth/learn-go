package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("result is: ", result)

	proResult := proAdder(2,5,7,9,11)
	fmt.Println("proResult is: ", proResult)
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func greeter() {
	fmt.Println("Namaste from golang")
}