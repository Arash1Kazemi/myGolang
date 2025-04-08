package main

import "fmt"

func main() {

	arash := User{"Arash", "arash@yaho.com", true, 19}
	fmt.Println(arash)
	fmt.Printf("user is: %v\n\n", arash)
	fmt.Printf("beteer: %+v\n\n", arash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
