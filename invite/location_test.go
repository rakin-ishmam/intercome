package invite

import (
	"math"
	"testing"
)

const TOLERANCE = 0.400001

func TestKiloDist(t *testing.T) {
	tt := []struct {
		name string
		src  Location
		dst  Location
		exp  float64
	}{
		{
			"dakha to cox's bazar",
			Location{number(23.8103), number(90.4125)},
			Location{number(21.4272), number(92.0058)},
			311.37,
		},
		{
			"dakha to dublin",
			Location{number(23.8103), number(90.4125)},
			Location{number(53.3498), number(6.2603)},
			7526.98,
		},
	}

	for _, tab := range tt {
		t.Run(tab.name, func(t *testing.T) {
			dist := tab.src.KiloDist(tab.dst)
			if diff := math.Abs(dist - tab.exp); diff > TOLERANCE { // consider equal and break
				t.Errorf("expected value is %v, but got %v", tab.exp, dist)
			}
		})

	}
}
