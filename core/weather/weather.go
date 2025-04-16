package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Name    string        `json:"name"`
	Main    MainData      `json:"main"`
	Weather []WeatherData `json:"weather"`
}

// "main": {
// 	"temp": 301.32,
// 	"feels_like": 305.04,
// 	"temp_min": 301.32,
// 	"temp_max": 301.32,
// 	"pressure": 1012,
// 	"humidity": 76,
// 	"sea_level": 1012,
// 	"grnd_level": 1011
//   },

type MainData struct {
	Temp     float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
}

// {
// 	"id": 801,
// 	"main": "Clouds",
// 	"description": "few clouds",
// 	"icon": "02d"
//   }

type WeatherData struct {
	Description string `json:"description"`
}

func FetchWeather() *WeatherResponse {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("WEATHER_API_KEY")
	CITY := os.Getenv("CITY")

	if API_KEY == "" || CITY == "" {
		log.Fatal("Missing API key or city in environment variables")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", CITY, API_KEY)
	// fmt.Println(url)
	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch weather: %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch weather data: %s", resp.Status)
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return &data
}
