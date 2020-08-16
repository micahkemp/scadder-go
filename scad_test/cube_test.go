package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestCube(t *testing.T) {
	tests := []struct {
		x, y, z float64
		center	bool
		n       string
		want    string
	}{
		{1, 2, 3, false, "", `module cube_component() {
	cube(center=false, size=[1, 2, 3]);
}

// call module when loaded directly
cube_component();
`},
		{1, 2, 3, true, "my_cube", `module my_cube() {
	cube(center=true, size=[1, 2, 3]);
}

// call module when loaded directly
my_cube();
`},
	}

	for _, test := range tests {
		c := scad.NewCube(test.x, test.y, test.z, test.center, test.n)

		if c.String() != test.want {
			t.Errorf("NewCube(%f, %f, %f, %t, %s).String() = %s, want %s",
				test.x,
				test.y,
				test.z,
				test.center,
				test.n,
				c.String(),
				test.want,
			)
		}
	}
}
