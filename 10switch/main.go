package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	num := rand.Intn(5)

	switch num {
	case 0:
		fmt.Println("aaa")
	case 1:
		fmt.Println("bbbb")
	case 2:
		fmt.Println("ccc")
		fallthrough
	case 3:
		fmt.Println("ddd")
		fallthrough
	case 4:
		fmt.Println("eeee")
	}
}
