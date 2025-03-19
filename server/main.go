package main

import (
	"fmt"
	"io"
	"html"
	"net/http"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}  
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func getWeatherReport(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
	api_key:=goDotEnvVariable("API_KEY")
	fmt.Println(api_key)
	// params:=r.URL.Query()
	
}

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/getWeatherReport",getWeatherReport)
	port := ":5000"
	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(port, nil) // A blocking call as it will listen only once executed
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
