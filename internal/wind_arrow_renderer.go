package internal

import "math"

func renderWindArrow(angle float64) string {
	angle = math.Mod(angle, 360)
	switch {
	case angle >= 337.5 || angle < 22.5:
		return "↑" // North
	case angle < 67.5:
		return "↗" // Northeast
	case angle < 112.5:
		return "→" // East
	case angle < 157.5:
		return "↘" // Southeast
	case angle < 202.5:
		return "↓" // South
	case angle < 247.5:
		return "↙" // Southwest
	case angle < 292.5:
		return "←" // West
	case angle >= 292.5:
		return "↖" // Northwest
	default:
		return "?" // Should never happen
	}
}
