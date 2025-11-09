package internal

import (
	"bytes"
	"html/template"
)

var tpl, parseErr = template.New("mail").Parse(`<h1>Wind alert!</h1>
<p>Location: {{printf "%.6f,%.6f" .Lat .Lon}}</p>
`)
func init() {
	if parseErr != nil {
		panic(parseErr)
	}
}

func RenderMail(weatherReading *WeatherReading) (string, error) {
	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, weatherReading)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
