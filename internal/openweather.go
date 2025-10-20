package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type reading struct {
	Dt        int     `json:"dt"`
	WindSpeed float64 `json:"wind_speed"`
	WindDeg   int     `json:"wind_deg"`
}

type OpenWeatherResponse struct {
	Lat    float64   `json:"lat"`
	Lon    float64   `json:"lon"`
	Hourly []reading `json:"hourly"`
	Daily  []reading `json:"daily"`
}

func CallOpenWeather(latitude float32, longitude float32, token string) (*WeatherReading, error) {
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
	var resp OpenWeatherResponse
	err := json.Unmarshal(*content, &resp)
	if err != nil {
		return nil, err
	}

	return &WeatherReading{
		Lat: resp.Lat,
		Lon: resp.Lon,
		Readings: map[string]*[]WindDataPoint{
			"hourly": getWindDatapoints(&resp.Hourly),
			"daily":  getWindDatapoints(&resp.Daily),
		},
	}, nil
}

func getWindDatapoints(readings *[]reading) *[]WindDataPoint {
	result := make([]WindDataPoint, len(*readings))
	for idx, reading := range *readings {
		result[idx] = WindDataPoint{
			Time:      time.Unix(int64(reading.Dt), 0),
			WindSpeed: float32(reading.WindSpeed),
			WindAngle: float32(reading.WindDeg),
		}
	}
	return &result
}
