package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"reflect"
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
func getCoordinates(city string) (float64, float64) {
	api_key := goDotEnvVariable("API_KEY")
	api := goDotEnvVariable("GEOENCODING_API")
	url := api+"?q="+city+"&limit=1&appid="+api_key
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-Key", api_key)
    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return -1, -1
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	//fmt.Println(body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return -1, -1
	}
	fmt.Printf("my type is: ")
	fmt.Println( reflect.TypeOf(body))
	latValue := 51.50
	lonValue := -0.12 
	return latValue, lonValue
}
func getWeatherReport(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getWeatherReport" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		//fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
	} else {
		http.Error(w, "Invalid request method.", 405)
		return 
	}
	api_key:=goDotEnvVariable("API_KEY")
	// fmt.Println(api_key)
	// fmt.Println(api_key)
	params:=r.URL.Query()
	api := goDotEnvVariable("WEATHER_API")
	// fmt.Println(api)
	city := params.Get("city")
	lat, lon := getCoordinates(city)
	fmt.Println(lat)
	fmt.Println(lon)
	latValue := strconv.FormatFloat(lat, 'E', -1, 64)
	lonValue := strconv.FormatFloat(lon, 'E', -1, 64)
	url := api+"?lat="+latValue+"&lon="+lonValue+"&appid="+api_key
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-Key", api_key)
    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("body in weather report " + string(body))
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
