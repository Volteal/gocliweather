package utilities

import (
	"fmt"
	"time"

	"github.com/Volteal/GoCliWeather/internal/models"
	"github.com/fatih/color"
)

func WeatherFormatter(m models.Weather) {
	location, current, hours := m.Location, m.Current, m.Forecast.Forecastday[0].Hour
	date := time.Unix(location.LocalTEpoch, 0)

	fmt.Printf(
		"Weather Forecast For: %s.\nLocation: %s, %s\nCurrent Weather: %.0f\u00B0C (%.0f\u00B0F), %s\n",
		date.Format("January 2, 2006"), location.Name, location.Country, current.TempC, current.TempF, current.Condition.Text,
	)

	for _, hour := range hours {
		dateTime := time.Unix(hour.TimeEpoch, 0)

		if dateTime.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf(
			"Time: %s - Temp: %.0f\u00B0C (%.0f\u00B0F), Chance of Rain: %.0f%%, Condition: %s\n",
			dateTime.Format("15:04"), hour.TempC, hour.TempF, hour.ChanceOfRain, hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
