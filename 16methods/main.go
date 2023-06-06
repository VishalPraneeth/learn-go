package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")

	vishal := User{"Vishal", "vishal@go.dev", 21, true}
	// fmt.Println(vishal)
	// fmt.Printf("Vishal details are: %+v\n", vishal)
	vishal.GetStatus()
	vishal.NewMail()
}

type User struct {
	Name string
	Email string
    Age int
	Status bool

}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)

}