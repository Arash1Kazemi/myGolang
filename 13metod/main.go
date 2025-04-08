package main

import "fmt"

func main() {

	arash := User{"Arash", "arash@yaho.com", true, 19}
	fmt.Println(arash)
	arash.pa()
	fmt.Println(arash.Name)
}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (A User) pa() {
	A.Name = "aaa"
	fmt.Println(A.Name)
}
