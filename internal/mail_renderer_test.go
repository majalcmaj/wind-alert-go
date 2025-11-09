package internal

import (
	"strings"
	"testing"
)

func TestRenderingMailDisplaysTitle(t *testing.T) {
	result, err := RenderMail(&WeatherReading{
		Lat:      40.7128,
		Lon:      -74.0060,
		Readings: map[string]*[]WindDataPoint{},
	})

	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if !strings.Contains(result, "Wind alert!") {
		t.Errorf("Expected 'Wind alert!' title to be present but got '%s'", result)
	}
}

func TestRenderingLatLonInformation(t *testing.T) {
	result, err := RenderMail(&WeatherReading{
		Lat:      40.72137,
		Lon:      -74.15497,
		Readings: map[string]*[]WindDataPoint{},
	})

	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if !strings.Contains(result, "40.721370,-74.154970") {
		t.Errorf("Expected lat/lon information to be present but got '%s'", result)
	}
}
