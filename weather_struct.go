package main

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float32 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		}`json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour [] struct {
				TimeEpoch int32 `json:"time_epoch"`
				TempC float32 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				}`json:"condition"`
				ChanceOfRain float32 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}