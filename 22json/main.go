package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	youCourse := []course{
		{"Golang", 300, "youtube", "go989ihj", []string{}},
		{"Python", 400, "youtube", "pyt989ihj", []string{"web", "dev", "python"}},
		{"Java", 500, "youtube", "jav989ihj", []string{"web", "dev", "java"}},
	}

	finalJson, err := json.MarshalIndent(youCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(finalJson)
	fmt.Printf("\n\n\n%s", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "golang",
		"price": 222,
		"website": "youtub",
		"tags": ["web-dev","backend-dev"]  
	}
	`)

	var T course

	checkValide := json.Valid(jsonDataFromWeb)

	if checkValide {
		fmt.Println("json is: ")
		json.Unmarshal(jsonDataFromWeb, &T)
		fmt.Printf("%#v\n\n\n", T)
	}

	var myOnloneData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnloneData)
	fmt.Printf("%#v\n\n", myOnloneData)

	for kr, valu := range myOnloneData {
		fmt.Printf("key is: %#v  value is: %#v type is: %T\n\n", kr, valu, valu)
	}
}
