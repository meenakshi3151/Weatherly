package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type ApiResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Load environment variables
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getCoordinates(city string) (float64, float64) {
	api_key := goDotEnvVariable("API_KEY")
	api := goDotEnvVariable("GEOENCODING_API")
	url := api + "?q=" + city + "&limit=1&appid=" + api_key
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
	var apiResponse ApiResponse
	response := string(body)
	fmt.Println(response)
	response = strings.TrimPrefix(response, "[")
	response = strings.TrimSuffix(response, "]")
	err = json.Unmarshal([]byte(response), &apiResponse)
	if err != nil {
		log.Fatal(err)
	}
	latValue := apiResponse.Lat
	lonValue := apiResponse.Lon
	fmt.Println(latValue)
	fmt.Println(lonValue)
	return latValue, lonValue
}

func getWeatherReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	api_key := goDotEnvVariable("API_KEY")
	params := r.URL.Query()
	api := goDotEnvVariable("WEATHER_API")
	city := params.Get("city")

	lat, lon := getCoordinates(city)
	if lat == -1 || lon == -1 {
		http.Error(w, "Error fetching coordinates", http.StatusInternalServerError)
		return
	}

	

	latValue := strconv.FormatFloat(lat, 'E', -1, 64)
	lonValue := strconv.FormatFloat(lon, 'E', -1, 64)
	url := api + "?lat=" + latValue + "&lon=" + lonValue + "&appid=" + api_key
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

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, fmt.Sprintf("Error reading response: %v", err), http.StatusInternalServerError)
		return
	}
    
}

func main() {
	godotenv.Load(".env")
	router := mux.NewRouter()

	router.HandleFunc("/", getRoot).Methods("GET")
	router.HandleFunc("/hello", getHello).Methods("GET")
	router.HandleFunc("/getWeatherReport", getWeatherReport).Methods("GET")

	// Setup CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow frontend
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Apply CORS middleware
	handler := c.Handler(router)

	port := ":5000"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
