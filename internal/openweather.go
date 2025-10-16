package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Current current `json:"current"`
}

type current struct {
	WindSpeed float32 `json:"wind_speed"`
	WindAngle float32 `json:"wind_deg"`
}

func CallOpenWeather(latitude float32, longitude float32, token string) (*WeatherReading, error){
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=current,minutely,alerts&appid=%s", latitude, longitude, token)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			_ = fmt.Errorf("warning: error when closing the response body: %+v", err)
		}
	}()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	// TODO: this is not good - the format is different.
	// See: https://openweathermap.org/api/one-call-3#example
	return ParseOpenweatherResponse(&body)
}

func ParseOpenweatherResponse(content *[]byte) (*WeatherReading, error) {
	var resp response
	err := json.Unmarshal(*content, &resp)
	if err != nil {
		return nil, err
	}

	return &WeatherReading{WindSpeed: resp.Current.WindSpeed, WindAngle: resp.Current.WindAngle}, nil
}
