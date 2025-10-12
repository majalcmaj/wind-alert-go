package internal

import (
	"os"
	"log"
	"testing"
	"reflect"
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

	expected := WeatherData{WindSpeed: 3.13, WindAngle: 93.0};
	if !reflect.DeepEqual(response, &expected) {
		t.Errorf("Expected parsed response was %+v but got %+v", &expected, content)
	}
}
