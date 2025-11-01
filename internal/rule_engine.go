package internal

type Range struct {
	From float64
	To   float64
}

type Rule struct {
	AngleRange Range
	SpeedRange Range
	HourRange  Range
}

func (rng Range) withinCyclicRange(angle float64) bool {
	if rng.From > rng.To {
		return angle >= rng.From || angle <= rng.To
	}
	return rng.withinRange(angle)
}

func (rng Range) withinRange(number float64) bool {
	return number >= rng.From && number <= rng.To
}

func RunRuleEngine(dataPoint WindDataPoint, rules *[]Rule) (bool, error) {
	for _, rule := range *rules {
		if rule.AngleRange.withinCyclicRange(dataPoint.WindAngle) &&
			rule.SpeedRange.withinRange(dataPoint.WindSpeed) &&
			rule.HourRange.withinCyclicRange(float64(dataPoint.Time.Hour())+float64(dataPoint.Time.Minute())/60.0) {
			return true, nil
		}
	}
	return false, nil
}
