package main

import (
	"DB/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is geting started ...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Server is running on port 8080 ...")
}