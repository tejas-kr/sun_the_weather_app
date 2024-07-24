package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"github.com/fatih/color"
)


func main() {
	var q string
	if len(os.Args) >= 2 {
		q = strings.Join(os.Args[1:], " ")
	} else {
		q = "new delhi"
	}
	fmt.Printf("Location: %s\n", q)

	raw_url := "http://api.weatherapi.com/v1/forecast.json?"
	processed_url := parse_url(raw_url, q)
	fmt.Println("Processed URL: ", processed_url)

	api_res := get_api_resp(processed_url)
	defer api_res.Body.Close()
	if api_res.StatusCode != 200 {
		panic("Weather API is not available!")
	}

	body, err := io.ReadAll(api_res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(int64(hour.TimeEpoch), 0)
		if date.Before(time.Now().UTC()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			color.Yellow(message)
		} else {
			color.Red(message)
		}
	}
	 
}
