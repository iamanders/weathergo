package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func main() {

	// Config file
	home := os.Getenv("HOME")
	configFileName := ".goweather"
	if !configFileExists(home, configFileName) {
		createBaseConfigFile(home, configFileName)
		fmt.Println("Hello! It seems like this is your first time starting goweather.")
		fmt.Println(fmt.Sprintf("A config file has been created at %s, edit the file and restart the program.", path.Join(home, configFileName)))
	}

	err := godotenv.Load(path.Join(home, configFileName))
	if err != nil {
		log.Fatal("Error loading config file")
	}

	// API key
	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("No OpenWeather API key (API_KEY) in config file")
	}

	// City
	city := os.Getenv("CITY")
	if city == "" {
		log.Fatal("No City (CITY) in config file")
	}

	// Get weather
	weather, err := GetWeather(key, city)
	if err != nil {
		log.Fatal(err)
	}

	// Print weather
	printWeather(&weather)
}
