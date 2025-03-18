package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
    http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	port := ":5000"
	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(port, nil) // A blocking call as it will listen only once executed 
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
	
}
