package main

import ("fmt"
		"encoding/json")


type course struct {
	Name string
	Price int
	Platform string
	Password string
	Tags []string
}
func main() {
	fmt.Println("Hello World!")
	//var myCourse course = course{Name: "Go", Price: 0, Platform: "Udemy"}
	//EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	learnCourse := []course{"Go", 299, "Udemy", "1234", []string{"development", "programming"},
		{"Python", 199, "Udemy", "1234", []string{"development", "programming"}},
		{"Java", 199, "Udemy", "1234", nil},
	}

	finalJson, err := json.MarshalIndent(learnCourse,"", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJson(){
	jsonString := []byte(`
	{
		"Name": "Go",
		"Price": 0,
		"Platform": "Udemy",
		"Password": "1234",
		"Tags": ["Go", "Programming", "Development"]
	}`)
	var myCourse course
	checkJsonString := json.Valid(jsonString)
	if checkJsonString{
		fmt.Println("This is a valid JSON")
		json.Unmarshal([]byte(jsonString), &myCourse)
		fmt.Printf("%+v\n", myCourse)
	}else{
		fmt.Println("This is not a valid JSON")
	}
	
	
}

// This code does not work for encoding😅😅