package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday" , "Saturday"}

	fmt.Println(days)

	// for d:= 0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Println(day)
    // }

	roughValue := 4
	for roughValue < 10 {
		fmt.Println(roughValue)
        roughValue++
	}

}