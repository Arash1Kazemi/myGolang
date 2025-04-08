package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang ")

	myCh := make(chan int,28)
	wg := &sync.WaitGroup{}


	wg.Add(2)

	go func (ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <- myCh  
		fmt.Println(val)
		fmt.Println(isChanelOpen)
		fmt.Println(<-myCh)

		wg.Done()
	}(myCh,wg)

	go func (ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 7
		myCh <- 5
		close(myCh)
		wg.Done()
	}(myCh,wg)

	wg.Wait()
	

}
