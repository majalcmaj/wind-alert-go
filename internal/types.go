package internal

import "time"

type WeatherReading struct {
	Lat      float64
	Lon      float64
	Readings map[string]*[]WindDataPoint
}

type WindDataPoint struct {
	Time      time.Time
	WindSpeed float32
	WindAngle float32
}
