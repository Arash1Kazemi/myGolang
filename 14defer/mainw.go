package main

import "fmt"

func main() {
	fmt.Println("1")
	defer first()
	defer fmt.Println("2")
}

func first() {
	fmt.Println("hello")
}

