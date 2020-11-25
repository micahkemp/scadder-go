package scad

import "github.com/micahkemp/scad/scad/internal"

type Rotate struct {
	scadTemplate
}

func NewRotate(name string, a, x, y, z float64, c ...DirPathRenderer) Rotate {
	n, _ := internal.FirstNonEmptyName(name, "rotate_component")
	fields := internal.NewFields(map[string]string{
		"a": internal.ShortFloat(a),
		"v": internal.ShortFloatList(x, y, z),
	})

	return Rotate{
		newTemplate(n, internal.TransformationTemplate, "rotate", fields, c...),
	}
}

func (t scadTemplate) Rotate(name string, a, x, y, z float64) Rotate {
	return NewRotate(name, a, x, y, z, t)
}
