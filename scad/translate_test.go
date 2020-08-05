package scad

import "testing"

func TestTranslate(t *testing.T) {
	child1 := NewCube(1, 2, 3, "")

	test := struct {
		x, y, z float64
		n       string
		c       []Template
		want    string
	}{
		1, 2, 3, "", []Template{child1}, `import <cube_component.scad>

module translate_module {
	translate(v=[1, 2, 3]) {
		cube_component();
	}
}`,
	}

	c := NewTranslate(test.x, test.y, test.z, test.n, child1)

	if c.String() != test.want {
		t.Errorf("NewTranslate(%f, %f, %f, %s, %v).String() = %s, want %s",
			test.x,
			test.y,
			test.z,
			test.n,
			test.c,
			c.String(),
			test.want,
		)
	}
}
