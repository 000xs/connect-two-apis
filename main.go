package main

import (
	"fmt"

	"github.com/000xs/connect-two-apis/core/weather"
	"github.com/000xs/connect-two-apis/core/webhook"
)

func main() {

	data := *weather.FetchWeather()
	fmt.Printf("🌤️ Weather in %s: %.1f°C, %s\n", data.Name, data.Main.Temp, data.Weather[0].Description)
	webhook.SendToWebhook(&data)
}
