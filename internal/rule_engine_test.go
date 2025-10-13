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
		t.Errorf("Rule engine should return false when no rules are passed")
	}
}

func TestPassingNoMatchingRules(t *testing.T) {
	reading := WeatherReading{WindSpeed: 8.0, WindAngle: 20.1}
	rules := []Rule{
		{ // Not matching anything
			SpeedRange: Range{From: 0.0, To: 7.0},
			AngleRange: Range{From: 95.0, To: 130.37},
		},
		{ // Not matching speed
			SpeedRange: Range{From: 20.1, To: 30.4},
			AngleRange: Range{From: 15.0, To: 21.37},
		},
		{ // Not matching angle
			SpeedRange: Range{From: 7.0, To: 13.5},
			AngleRange: Range{From: 0.2, To: 2.137},
		},
	}

	result, err := RunRuleEngine(reading, &rules)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != false {
		t.Errorf("Rule engine should return false for no matching rules")
	}
}

func TestPassingSingleMatchingRule(t *testing.T) {
	reading := WeatherReading{WindSpeed: 8.0, WindAngle: 20.1}
	rules := []Rule{
		{ // Not matching rule
			SpeedRange: Range{From: 0.0, To: 7.0},
			AngleRange: Range{From: 95.0, To: 130.37},
		},
		{ // Finally - a matching rule
			SpeedRange: Range{From: 0.0, To: 10.0},
			AngleRange: Range{From: 15.0, To: 21.37},
		},
		{ // Another not matching angle
			SpeedRange: Range{From: 7.0, To: 13.5},
			AngleRange: Range{From: 0.2, To: 2.137},
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

func TestRuleWithAngleRangeFromHigherThanToAngleLessThanTo(t *testing.T) {
	reading := WeatherReading{WindSpeed: 5.0, WindAngle: 83}
	rules := []Rule{
		{
			SpeedRange: Range{From: 0.0, To: 7.0},
			AngleRange: Range{From: 270.0, To: 90.0},
		},
	}

	result, err := RunRuleEngine(reading, &rules)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != true {
		t.Errorf("Rule engine should return true when rule's from is higher than to and angle is less than to")
	}
}

func TestRuleWithAngleRangeFromHigherThanToAngleBiggerThanFrom(t *testing.T) {
	reading := WeatherReading{WindSpeed: 5.0, WindAngle: 300}
	rules := []Rule{
		{
			SpeedRange: Range{From: 0.0, To: 7.0},
			AngleRange: Range{From: 270.0, To: 90.0},
		},
	}

	result, err := RunRuleEngine(reading, &rules)

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
	if result != true {
		t.Errorf("Rule engine should return true when rule's from is higher than to and angle is bigger than from")
	}
}
