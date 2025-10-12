package internal

import (
	"testing"
)

func TestPassingEmptyRule(t *testing.T) {
	reading := WeatherReading{WindSpeed: 10.0, WindAngle: 30.0}
	rules := []Rule{}

	result, err := RunRuleEngine(reading, &rules)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != false {
		t.Errorf("Rule engine should return false for empty rules")
	}
}

func TestPassingSingleMatchingRule(t *testing.T) {
	reading := WeatherReading{WindSpeed: 8.0, WindAngle: 20.1}
	rules := []Rule{
		{
			SpeedRange: Range{From: 0.0, To: 10.0},
			AngleRange: Range{From: 15.0, To: 21.37},
		},
	}

	result, err := RunRuleEngine(reading, &rules)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != true {
		t.Errorf("Rule engine should return true for a data matching the rule")
	}
}
