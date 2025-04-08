package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//performGetRequest()
	// performPostRequest()
	performPostFormRequest()
}

func performGetRequest() {

	const myuUrl = "http://localhost:8000/get"

	respon, err := http.Get(myuUrl)
	if err != nil {
		panic(err)
	}

	// fmt.Println(respon.Body)
	// fmt.Println(respon.ContentLength)

	// content, err := io.ReadAll(respon.Body)

	// // fmt.Println(string(content))
	defer respon.Body.Close()

	var responseString strings.Builder
	content, _ := io.ReadAll(respon.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}

func performPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"lange":"golang"
			"lerner":"ash"
			"platform":"youtub"
		}
	`)
	respons, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer respons.Body.Close()
	content, err := io.ReadAll(requestBody)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	fmt.Println("respons is: ", respons)
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("name", "ash")
	data.Add("lastname", "kz")

	respon, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(respon.Body)
	fmt.Println(content)
	defer respon.Body.Close()
	
}
