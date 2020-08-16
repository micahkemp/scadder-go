package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestCylinderByRadius(t *testing.T) {
	tests := []struct {
		r, h   float64
		center bool
		n      string
		want   string
	}{
		{1, 2, false, "", `module cylinder_component() {
	cylinder(center=false, h=2, r=1);
}

// call module when loaded directly
cylinder_component();
`},
	}

	for _, test := range tests {
		c := scad.NewCylinderByRadius(test.r, test.h, test.center, test.n)

		if c.String() != test.want {
			t.Errorf("NewCylinderByRadius(%f, %f, %t, %s).String() = %s, want %s",
				test.r,
				test.h,
				test.center,
				test.n,
				c.String(),
				test.want,
			)
		}
	}
}
