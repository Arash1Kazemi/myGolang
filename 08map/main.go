package main

import "fmt"

func main() {

	game := make(map[string]string)

	game["d2"] = "dota2"
	game["hsd"] = "huntShowDown"
	game["sf"] = "splitFiction"
	fmt.Println(game["d2"])
	delete(game, "d2")
	fmt.Println(game)

	//
	for _, v := range game {
		fmt.Printf("For key %v Value is %v\n", v)

	}

}
