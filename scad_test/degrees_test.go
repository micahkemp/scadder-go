package scad

import (
	"github.com/micahkemp/scad/scad"
	"math"
	"testing"
)

func TestDegrees(t *testing.T) {
	tests := []struct {
		degrees float64
		radians float64
	}{
		{0, 0},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, math.Pi * 1.5},
		{360, math.Pi * 2},
	}

	for _, test := range tests {
		if scad.DegreesToRadians(test.degrees) != test.radians {
			t.Errorf("DegreesToRadians(%f) = %f, want %f",
				test.degrees,
				scad.DegreesToRadians(test.degrees),
				test.radians)
		}

		if scad.RadiansToDegrees(test.radians) != test.degrees {
			t.Errorf("RadiansToDegrees(%f) = %f, want %f",
				test.radians,
				scad.RadiansToDegrees(test.radians),
				test.degrees)
		}
	}
}
