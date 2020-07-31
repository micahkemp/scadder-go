package solids

import (
	"github.com/micahkemp/scad/component"
	"github.com/micahkemp/scad/renderer"
	"testing"
)

func TestFields(t *testing.T) {
	tests := []struct {
		l, w, h float64
		want    map[string]string
	}{
		// int-like floats
		{1, 2, 3, map[string]string{"size": "[1, 2, 3]"}},
		// low precision floats
		{1.1, 2.2, 3.3, map[string]string{"size": "[1.1, 2.2, 3.3]"}},
	}

	for _, test := range tests {
		c := Cube{test.l, test.w, test.h, renderer.Renderer{}}
		f := c.Fields()
		ok := true
		if len(f) != len(test.want) {
			ok = false
		}

		for k, v := range f {
			if v != test.want[k] {
				ok = false
			}
		}

		if !ok {
			t.Errorf("%v.Fields() = %v, want %v", c, f, test.want)
		}
	}
}

func TestInterface(t *testing.T) {
	component.RenderFieldsComponent(Cube{})
}
