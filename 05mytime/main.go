package main
 import ("fmt"	
 		"time")

 func main() {
	fmt.Println("Welcome to time study of GoLang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 27, 9, 23, 0, 0, time.UTC)
	fmt.Println("Created Date:", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
 }