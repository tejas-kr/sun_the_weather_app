package main

import (
	"net/url"
	"net/http"
)

func parse_url(url_string string, city string) (string) {
	params := url.Values{}
	params.Add("key", "931a095960b84c09ae341358242207")
	params.Add("q", city)
	params.Add("days", "1")
	params.Add("aqi", "no")
	params.Add("alerts", "no")

	return url_string + params.Encode()
}


func get_api_resp(processed_url string) (http.Response) {
	api_res, err := http.Get(processed_url)
	if err != nil {
		panic(err)
	}
	return *api_res
}
