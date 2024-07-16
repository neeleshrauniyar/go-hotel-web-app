package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name       string   `json:"name"`
		Age        int      `json:"age"`
		IsStudent  bool     `json:"is_student"`
		EnrolledIn []string `'json:"enrolledIn"`
	}

	var p = Person{
		Name:      "james",
		Age:       18,
		IsStudent: true,
	}

	fmt.Println(p)

	jsonData, err := json.Marshal(p) //convert data type to a json using marshal method under json module

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonData)         //raw json data that is not understandable
	fmt.Println(string(jsonData)) //convert json data to json format

	myJSON := `{
		"name": "andrew", 
		"age": 30, 
		"is_student": true,
		"enrolledIn": ["math", "english"]
		}`

	fmt.Println(myJSON)

	var p2 Person

	//Concept: While unmarshalling the data, the data needs to be stored in a struct
	//that struct should have the fields that are present in the json data
	//else the missing fields will be ignored
	//the unmarshal method is made in such a way as below

	json.Unmarshal([]byte(myJSON), &p2) //convert json data to struct using unmarshal method under json module, json.Unmarshal([]byte(myJSON), &p)

	fmt.Println(p2)
}
