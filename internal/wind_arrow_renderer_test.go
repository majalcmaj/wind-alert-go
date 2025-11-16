package internal

import (
	"testing"
)

func TestRendersWindArrowsCorrectly(t *testing.T) {
	input := map[float64]string{
		338: "↑",
		372: "↑",
		0:   "↑",
		21:  "↑",
		23:  "↗",
		45:  "↗",
		67:  "↗",
		68:  "→",
		90:  "→",
		112: "→",
		113: "↘",
		135: "↘",
		157: "↘",
		180: "↓",
		202: "↓",
		203: "↙",
		225: "↙",
		247: "↙",
		270: "←",
		292: "←",
		293: "↖",
		315: "↖",
		337: "↖",
	}

	for angle, expectedRune := range input {
		result := renderWindArrow(angle)
		if result != expectedRune {
			t.Errorf("For angle %f expected %s but got %s", angle, expectedRune, result)
		}
	}
}
