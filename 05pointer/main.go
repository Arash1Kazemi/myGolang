package main

import "fmt"

func main() {

	var poi *int
	fmt.Println(" i ", poi)
	i0 := 23
	var I1 = &i0;
	fmt.Println(i0)
	fmt.Println(&i0)
	fmt.Println(I1)
	fmt.Println(&I1)
	fmt.Println(*I1)

}
