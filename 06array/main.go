package main

import "fmt"

func main() {
	var myGame [100]string

	myGame[0] = "cs2"
	myGame[1] = "dota"
	myGame[99] = "hunt"
	fmt.Println(myGame) 

	var veg  = [5]string{"RR","EE"}
	fmt.Println(veg)
}
