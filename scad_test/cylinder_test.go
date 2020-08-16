package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestCylinderByRadius(t *testing.T) {
	tests := []struct {
		n      string
		r, h   float64
		center bool
		want   string
	}{
		{"", 1, 2, false, `module cylinder_component() {
	cylinder(center=false, h=2, r=1);
}

// call module when loaded directly
cylinder_component();
`},
	}

	for _, test := range tests {
		c := scad.NewCylinderByRadius(test.n, test.r, test.h, test.center)

		if c.String() != test.want {
			t.Errorf("NewCylinderByRadius(%q, %f, %f, %t).String() = %s, want %s",
				test.n,
				test.r,
				test.h,
				test.center,
				c.String(),
				test.want,
			)
		}
	}
}
