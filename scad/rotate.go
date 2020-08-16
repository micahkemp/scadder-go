package scad

import "github.com/micahkemp/scad/scad/internal"

func NewRotate(n string, a, x, y, z float64, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "rotate_component")
	fields := internal.NewFields(map[string]string{
		"a": internal.ShortFloat(a),
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "rotate", fields, c...)
}

func (t scadTemplate) Rotate(n string, a, x, y, z float64) scadTemplate {
	return NewRotate(n, a, x, y, z, t)
}
