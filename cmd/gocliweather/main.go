package main

import (
	"fmt"
	"os"

	"github.com/Volteal/GoCliWeather/internal/weather"
)

func main() {
	var query string

	if len(os.Args) != 2 {
		fmt.Println("Error: this command requires an argument.")
		fmt.Println("Usage: gocliweather <location>")
		fmt.Println("Example: gocliweather London")
		return
	} else {
		query = os.Args[1]
	}

	weather.GetWeather(query)
}
