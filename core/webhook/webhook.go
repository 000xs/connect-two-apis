package webhook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/000xs/connect-two-apis/core/weather"
	"github.com/joho/godotenv"
)

// var data weather.WeatherResponse

func SendToWebhook(data *weather.WeatherResponse) {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//get webhook api and Url from .env
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	//cheack WEBHOOK_URL

	if WEBHOOK_URL == " " {
		panic("WEBHOOK_URL missing !")

	}

	// Encode data to JSON
	payload, err := json.Marshal(*data)
	if err != nil {
		log.Fatalf("Failed to encode payload: %v", err)
	}

	req, err := http.NewRequest("POST", WEBHOOK_URL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{Timeout: time.Minute}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	log.Println("✅ Weather data sent to webhook!")

}
