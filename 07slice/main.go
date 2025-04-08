package main

import "fmt"

func main() {

	var games = []string{}
	fmt.Printf("type of the game tte is: %T\n", games)

	games = append(games, "cs", "spliteFiction")
	fmt.Println(games)

	games = append(games[:1])
	fmt.Println(games)
	
	//how to remove 

	var course = []string{"Go", "rust", "python", "java"}
	fmt.Println(course)
	int := 2;
	fmt.Println()
	course = append(course[:int], course[int+1:]...)
	fmt.Println(course)
}
