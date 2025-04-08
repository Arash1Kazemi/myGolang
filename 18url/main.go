package main

import (
	"fmt"
	"net/url"
)

func main() {
	const myurl string = "https://loc.dev:8000/learn?course=go&lesson=18url"
	fmt.Println("myurl:", myurl)
// 	resalt, err := url.Parse(myurl)
// 	if err != nil {
// 		panic(err)
// 	}
//  qparam := resalt.Query()

	// for _, value := range qparam {

	// 	fmt.Println(value)
	// }

	pareOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "tutcss",
		RawPath: " user=arash",
	}
	p := pareOfUrl.String()	
	fmt.Println("pareOfUrl:", p)
}
