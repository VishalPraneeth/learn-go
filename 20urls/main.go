package main

import ("fmt"
		"net/url")

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=zxcvbnnn"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of the query parameters are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=vishal",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}