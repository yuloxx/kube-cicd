package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	fmt.Println("Server started at http://localhost:8080")
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
