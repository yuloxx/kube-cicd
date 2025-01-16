package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	myenv := os.Getenv("EXAMPLE_ENV")
	log.Println("EXAMPLE_ENV:", myenv)
	fmt.Fprintln(w, "hello: ", myenv)
}

func main() {
	myenv := os.Getenv("EXAMPLE_ENV")
	log.Println("EXAMPLE_ENV:", myenv)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
