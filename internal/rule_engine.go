package internal

type Range struct {
	From float32
	To   float32
}

type Rule struct {
	AngleRange Range
	SpeedRange Range
}

func (rng Range) withinAngleRange(angle float32) bool {
	if rng.From > rng.To {
		return angle >= rng.From || angle <= rng.To
	}
	return rng.withinRange(angle)
}

func (rng Range) withinRange(number float32) bool {
	return number >= rng.From && number <= rng.To
}

func RunRuleEngine(reading WeatherReading, rules *[]Rule) (bool, error) {
	for _, rule := range *rules {
		if rule.AngleRange.withinAngleRange(reading.WindAngle) &&
			rule.SpeedRange.withinRange(reading.WindSpeed) {
			return true, nil
		}
	}
	return false, nil
}
