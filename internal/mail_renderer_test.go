package internal

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"
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

	if !strings.Contains(result, "Wind Alert!") {
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

func TestRenderingDailyAndHourlyTables(t *testing.T) {
	reading := WeatherReading{
		Lat: 40.72137,
		Lon: -74.15497,
		Readings: map[string]*[]WindDataPoint{
			"daily": {
				{Time: parseTime("2025-01-01T10:00"), WindSpeed: 10, WindAngle: 180},
				{Time: parseTime("2025-01-02T10:00"), WindSpeed: 20, WindAngle: 80},
			},
			"hourly": {
				{Time: parseTime("2025-01-01T10:00"), WindSpeed: 10, WindAngle: 180},
				{Time: parseTime("2025-01-01T11:00"), WindSpeed: 12, WindAngle: 190},
			},
		},
	}

	renderedMail, err := RenderMail(&reading)

	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	for _, row := range *reading.Readings["daily"] {
		matchRow(t, renderedMail, row)
	}

	for _, row := range *reading.Readings["hourly"] {
		matchRow(t, renderedMail, row)
	}
}

func parseTime(tStr string) time.Time {
	tm, err := time.Parse(time.RFC3339, tStr+":00Z")
	if err != nil {
		panic(err)
	}
	return tm
}

func matchRow(t *testing.T, mailHtml string, row WindDataPoint) {
	tableRegEx := regexp.MustCompile(
		fmt.Sprintf(`(?s)<tr.*>.*<td.*>%s</td>.*<td.*>%.1f</td>.*<td.*>%s</td>.*</tr>`, row.Time.Format("2006-01-02 15:04"), row.WindSpeed, renderWindArrow(row.WindAngle)))

	match := tableRegEx.MatchString(mailHtml)
	if !match {
		t.Errorf("Expected WindDataPoint %+v not found in:\n%s", row, mailHtml)
	}
}
