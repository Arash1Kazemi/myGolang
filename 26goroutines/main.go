package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal = []string{"hallo"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	website := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.reddit.com",
	}

	for _, web := range website {
		go getStstusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
}

func getStstusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("err in endpoint")
	} else {
		mut.Lock()
		signal = append(signal, endpoint)
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
		mut.Unlock()
	}
}
