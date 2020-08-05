package scad

import "testing"

func TestCube(t *testing.T) {
	tests := []struct {
		x, y, z float64
		n       string
		want    string
	}{
		{1, 2, 3, "", `module cube_component {
	cube(size=[1, 2, 3]);
}`},
	}

	for _, test := range tests {
		c := NewCube(test.x, test.y, test.z, test.n)

		if c.String() != test.want {
			t.Errorf("NewCube(%f, %f, %f, %s).String() = %s, want %s",
				test.x,
				test.y,
				test.z,
				test.n,
				c.String(),
				test.want,
			)
		}
	}
}
