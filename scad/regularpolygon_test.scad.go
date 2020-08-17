package scad

import (
	"testing"
)

func TestRegularPolygon(t *testing.T) {
	tests := []struct {
		attrs        regularPolygonAttributes
		pointsString string
	}{
		{
			regularPolygonAttributes{
				4,
				1,
			},
			"[[1, 0], [0, 1], [-1, 0], [0, -1]]",
		},
	}

	for _, test := range tests {
		if test.attrs.points().String() != test.pointsString {
			t.Errorf("%v.points().String() = %s, want %s",
				test.attrs,
				test.attrs.points().String(),
				test.pointsString)
		}
	}
}
