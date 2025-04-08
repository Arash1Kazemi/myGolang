package main

import (
	"fmt"
	"io"
	"net/http"
)

const urel = "https://api.github.com/users/"

func main() {
	respon, err := http.Get(urel)
	if err != nil {
		panic(err)
	}
	// fmt.Println(respon)
	// fmt.Println("Response status:", respon.Status)
	defer respon.Body.Close()
	databyte, err := io.ReadAll(respon.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println("Response content:", content)
	fmt.Println("Response length:", len(content))
}
