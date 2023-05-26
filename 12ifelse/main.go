package main

import "fmt"

func main()  {
	fmt.Println("If else in Golang")
	logincount := 10
	var result string
	if logincount < 10 {
		result = "regular user"
	}else if logincount>10{
		result = "Watch"
	}else{
        result = "Exactly 10"
    }
	fmt.Println("Result: ", result)

	if 9%2 == 0{
		fmt.Println("Number is even")

	} else{
		fmt.Println("Number is odd")
	}

	if num := 3; num<10 {
		fmt.Println(num, "is less than 10")
	} else{
		fmt.Println(num, "isn't less than 10")
	}

	// if err != nil
}