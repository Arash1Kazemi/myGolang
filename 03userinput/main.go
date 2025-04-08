package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Print("you rate us ", input)

	fmt.Printf("The Type Is: %T\n", input)

	rateAddOne, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("i want to add one to youre rate\n", rateAddOne+1)
	}
}
