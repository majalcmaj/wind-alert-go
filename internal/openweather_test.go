package internal

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func readTestData() (*[]byte, error) {
	content, err := os.ReadFile("testdata/openWeatherMap.json")
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func TestParsingOutValues(t *testing.T) {
	content, err := readTestData()

	if err != nil {
		log.Fatal(err)
	}

	response, err := ParseOpenweatherResponse(content)

	if err != nil {
		log.Fatal(err)
	}

	expected := WeatherReading{
		Lat: 33.44,
		Lon: -94.04,
		Readings: map[string]*[]WindDataPoint{
			"hourly": {
				{
					Time:      time.Unix(1684926000, 0),
					WindSpeed: 2.58,
					WindAngle: 86,
				},
			},
			"daily": {
				{
					Time:      time.Unix(1684951200, 0),
					WindSpeed: 3.98,
					WindAngle: 76,
				},
			},
		},
	}
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Expected parsed response was %+v but got %+v", &expected, response)
	}
}
