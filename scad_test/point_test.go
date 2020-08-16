package scad

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestPoint2D(t *testing.T) {
	tests := []struct {
		x    float64
		y    float64
		want string
	}{
		{1, 2, "[1, 2]"},
		{1.0, 2.1, "[1, 2.1]"},
	}

	for _, test := range tests {
		point := scad.Point2D{X: test.x, Y: test.y}
		if point.String() != test.want {
			t.Errorf("Point{X: %f, Y: %f}.String() = %s, want %s",
				test.x,
				test.y,
				point.String(),
				test.want)
		}
	}
}

func TestPoints(t *testing.T) {
	tests := []struct {
		points scad.Points2D
		want   string
	}{
		{
			scad.Points2D{{1, 2}, {3, 4}},
			"[[1, 2], [3, 4]]",
		},
	}

	for _, test := range tests {
		if test.points.String() != test.want {
			t.Errorf("%v.String() = %s, want %s",
				test.points,
				test.points.String(),
				test.want)
		}
	}
}
