package main

import "fmt"

func main() {
fmt.Println(sumPro(8,9,0,0,0))
}

func aa() {
	fmt.Println("1111")
}

func sum(a int, b int) int {
	return a + b
}

func sumPro(vari ...int) (string,int) {
	total := 0
	for _, v := range vari {
		total += v;
	}
	return "hallo",total
}
