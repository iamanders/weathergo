package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WeatherResponse struct {
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
	} `json:"city"`
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt      int `json:"dt"`
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
		Temp    struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		Weather  []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Speed  float64 `json:"speed"`
		Deg    int     `json:"deg"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Snow   float64 `json:"snow,omitempty"`
	} `json:"list"`
}

// Get weather from Open weathermap API
func GetWeather(key string, city string) (WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast/daily?q=%s&cnt=5&appid=%s&units=metric&lang=sv", city, key)

	httpClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return WeatherResponse{}, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return WeatherResponse{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	weather := WeatherResponse{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil
}
