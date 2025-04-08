package main

import (
	"fmt"

	"github.com/go-delve/delve/pkg/dwarf/reader"
)

func main() {

		input := reader.ReadString('\n');
		fmt.Println("rate us", input);
		fmt.Printf("The Type Is: %T", input);
}
