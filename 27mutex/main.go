package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition -Mutex")

	// Create a WaitGroup to wait for all goroutines to finish
	wg := &sync.WaitGroup{}
	// Create an RWMutex to protect shared data
	mut := &sync.RWMutex{}

	//mut.RLock()  Not Good Here
	var score = []int{0}
	//mut.RUnlock() Not Good Here

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock() // Lock the mutex before modifying shared data
		score = append(score, 1)
		m.Unlock() // Unlock the mutex after modifying shared data
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("score", score)
}
