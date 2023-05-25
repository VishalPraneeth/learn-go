package main

import "fmt"

func main()  {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["php"] = "php"
	languages["python"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Js shorts for: ", languages["js"])

	delete(languages, "php")
	fmt.Println("List of all languages: ", languages)

	// loops

	for key,value := range languages {
		fmt.Printf("Key: %v\n", key)
		fmt.Printf("Value: %v", value)
	}
}