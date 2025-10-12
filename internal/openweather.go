package internal

import (
	"encoding/json"
)

type response struct {
	Current current `json:"current"`
}

type current struct {
	WindSpeed float32 `json:"wind_speed"`
	WindAngle float32 `json:"wind_deg"`
}

func ParseOpenweatherResponse(content *[]byte) (*WeatherData, error) {
	var resp response
	err := json.Unmarshal(*content, &resp)
	if err != nil {
		return nil, err
	}
	
	return &WeatherData{WindSpeed: resp.Current.WindSpeed, WindAngle: resp.Current.WindAngle}, nil
}
