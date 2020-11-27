package scad

import (
	"testing"
)

func TestCube(t *testing.T) {
	tests := []struct {
		n       string
		x, y, z float64
		center  bool
		want    string
	}{
		{"", 1, 2, 3, false, `module cube_component() {
	cube(center=false, size=[1, 2, 3]);
}
// call module when loaded directly
cube_component();
`},
		{"my_cube", 1, 2, 3, true, `module my_cube() {
	cube(center=true, size=[1, 2, 3]);
}
// call module when loaded directly
my_cube();
`},
	}

	for _, test := range tests {
		c := NewCube(test.n, test.x, test.y, test.z, test.center)

		if c.ScadTemplate.String() != test.want {
			t.Errorf("NewCube(%s, %f, %f, %f, %t).String() = %s, want %s",
				test.n,
				test.x,
				test.y,
				test.z,
				test.center,
				c.ScadTemplate.String(),
				test.want,
			)
		}
	}
}
