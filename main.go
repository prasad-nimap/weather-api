package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Check if the city argument is provided
	if len(os.Args) < 2 {
		log.Fatal("Please provide the city name as an argument.")
	}

	// Set up the HTTP client
	client := &http.Client{}

	// Create a new GET request to the RapidAPI Weather API
	req, err := http.NewRequest("GET", "https://weather-api99.p.rapidapi.com/city-weather", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the required headers
	/*req.Header.Set("X-RapidAPI-Key", "YOUR_RAPIDAPI_KEY")
	req.Header.Set("X-RapidAPI-Host", "weather-api99.p.rapidapi.com")*/

	req.Header.Set("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")
	req.Header.Set("X-RapidAPI-Host", "weather-api99.p.rapidapi.com")

	// Get the city name from command-line arguments
	city := os.Args[1]

	// Add the city as a query parameter
	query := req.URL.Query()
	query.Add("city", city)
	req.URL.RawQuery = query.Encode()

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Print the weather data
	fmt.Println(string(body))
}
