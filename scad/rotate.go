package scad

import "github.com/micahkemp/scad/scad/internal"

func NewRotate(name string, a, x, y, z float64, c ...DirPathRenderer) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "rotate_component")
	fields := internal.NewFields(map[string]string{
		"a": internal.ShortFloat(a),
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(n, internal.TransformationTemplate, "rotate", fields, c...)
}

func (t scadTemplate) Rotate(name string, a, x, y, z float64) scadTemplate {
	return NewRotate(name, a, x, y, z, t)
}
