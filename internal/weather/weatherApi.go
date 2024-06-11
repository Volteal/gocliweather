package weather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Volteal/GoCliWeather/internal/models"
	"github.com/Volteal/GoCliWeather/internal/utilities"
)

func GetWeather(q string) {
	jsonData := fetchWeatherJson(q)
	handleFormatting(handleWeatherJson(jsonData))
}

func fetchWeatherJson(q string) *http.Response {
	data, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=a14d4f0137374e6b8ac144117241106&q=" + q + "&days=2&aqi=no&alerts=no")
	if err != nil {
		log.Fatal(data)
	}

	return data
}

func handleWeatherJson(r *http.Response) models.Weather {
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Fatal("Weather API not available.")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	return weather
}

func handleFormatting(m models.Weather) {
	utilities.WeatherFormatter(m)

}
