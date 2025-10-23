package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/majalcmaj/wind-alert-go/internal"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	openWeatherToken := os.Getenv("OPENWEATHER_TOKEN")

	if len(strings.TrimSpace(openWeatherToken)) == 0 {
		return nil, errors.New("OPENWEATHER_TOKEN env variable must be set")
	}

	openWeather, err := internal.NewOpenWeather("https://api.openweathermap.org", openWeatherToken)
	if err != nil {
		return nil, err
	}
	forecast, err := openWeather.GetForecast(54.646034, 18.512407)

	if err != nil {
		return nil, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%+v", forecast),
	}
	return &response, nil
}

func main() {
	lambda.Start(handler)
}
