package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiKey = "" //your apiKey
	apiURL = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherData struct {
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
}

func main() {
	var city string
	fmt.Println("Write a city")
	fmt.Scan(&city)
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", apiURL, city, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()

	var weatherData WeatherData

	err = json.NewDecoder(resp.Body).Decode(&weatherData)

	if err != nil {
		fmt.Println("decode err:", err)
		return
	}

	fmt.Printf("Current temp in %s: %.1f°C\n", city, weatherData.Main.Temp)
}
