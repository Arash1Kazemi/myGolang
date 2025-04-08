package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("see the peresent time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	datIWantToeIn:=time.Date(2000, time.December, 12299999999999,2444999999,23,23,0,time.UTC)
	fmt.Println(datIWantToeIn)
}
