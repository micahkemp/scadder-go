package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestPolygon(t *testing.T) {
	tests := []struct {
		n      string
		points scad.Points2D
		want   string
	}{
		{"", scad.Points2D{{1, 2}}, `module polygon_component() {
	polygon(points=[[1, 2]]);
}

// call module when loaded directly
polygon_component();
`},
		{"my_polygon", scad.Points2D{{1.1, 2.0}}, `module my_polygon() {
	polygon(points=[[1.1, 2]]);
}

// call module when loaded directly
my_polygon();
`},
	}

	for _, test := range tests {
		c := scad.NewPolygon(test.n, test.points)

		if c.String() != test.want {
			t.Errorf("NewCube(%s, %v).String() = %s, want %s",
				test.n,
				test.points,
				c.String(),
				test.want,
			)
		}
	}
}
