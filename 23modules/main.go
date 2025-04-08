package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("start")

	r := mux.NewRouter()
	r.HandleFunc("/", severHome)

	log.Fatal(http.ListenAndServe(":9000", r))
}

func severHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Wellcome to My Server </h1>"))
}
