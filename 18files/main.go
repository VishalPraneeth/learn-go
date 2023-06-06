package main

import ("fmt"
		"os"
		"io"
		"io/ioutil")

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myFile.txt") 
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is: \n", string(databyte))
}