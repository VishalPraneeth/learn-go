package main
 
import ("fmt"
		"sort")

func main() {
	fmt.Println("Welcome to video of slices in golang")

	var fruitList = []string{"Apple", "Orange", "Grapes"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 567
	//highScore[4] = 678

	highScore = append(highScore, 678, 789)

	fmt.Println(highScore)
	sort.Ints(highScore)
	// fmt.Println(highScore)
	// how to delete an element from slice based on index

	var courses = []string{"Docker", "Kubernetes", "Puppet", "Terraform", "AWS", "Azure", "GCP"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}