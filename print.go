package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Display weather
func printWeather(weather *WeatherResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, fmt.Sprintf("%s \tTemp C \tMin C \tMax C \tVäder \t", weather.City.Name))

	for i, day := range weather.List {
		dateString := "Idag"
		if i > 0 && i < 2 {
			dateString = "Imorgon"
		} else if i > 1 {
			dateString = strings.Title(strings.Repeat("över", i-1) + "morgon")
		}

		fmt.Fprintln(w, fmt.Sprintf("%s \t%.1f \t%.1f \t%.1f \t%s \t", dateString, day.Temp.Day, day.Temp.Min, day.Temp.Max, strings.Title(day.Weather[0].Description)))
	}

	w.Flush()
}
