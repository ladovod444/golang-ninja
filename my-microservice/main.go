package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func More(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "More")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/more", More)
	http.ListenAndServe(":8080", nil)
}
