package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"reflect"
)

type response struct {
	Current current `json:"current"`
}

type current struct {
	WindSpeed float32 `json:"wind_speed"`
	WindAngle float32 `json:"wind_deg"`
}

func readTestData() (*response, error) {
	file, err := ioutil.ReadFile("testdata/openWeatherMap.json")
	if err != nil {
		return nil, err
	}

	var resp response
	if json.Unmarshal(file, &resp) != nil {
		return nil, err
	}
	return &resp, nil
}

func TestParsingOutValues(t *testing.T) {
	testData, err := readTestData()

	if err != nil {
		log.Fatal(err)
	}

	expected := response{Current: current{WindSpeed: 3.13, WindAngle: 93.0}};
	if !reflect.DeepEqual(testData, &expected) {
		t.Errorf("Expected parsed response was %+v but got %+v", &expected, testData)
	}
}
