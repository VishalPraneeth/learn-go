package main

import ("fmt"
		"os"
		"bufio"
		"strconv"
		"strings")

func main() {
	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating our pizza", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else{
		fmt.Println("added 1 to your rating: ", numRating + 1)
	}
}