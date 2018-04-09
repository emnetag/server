package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Default)

	http.Handle("/", r)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
