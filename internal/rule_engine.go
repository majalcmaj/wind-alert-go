package internal

type Range struct {
	From float32
	To   float32
}

type Rule struct {
	AngleRange Range
	SpeedRange Range
}

func (_ Range) withinAngleRange(angle float32) bool {
	return true
}

func RunRuleEngine(reading WeatherReading, rules *[]Rule) (bool, error) {
	for _, rule := range *rules {
		if reading.WindAngle >= rule.AngleRange.From && reading.WindAngle <= rule.AngleRange.To &&
			reading.WindSpeed >= rule.SpeedRange.From && reading.WindSpeed <= rule.SpeedRange.To {
			return true, nil
		}
	}
	return false, nil
}
