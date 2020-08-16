package scad

import "github.com/micahkemp/scad/scad/internal"

func NewRotate(a, x, y, z float64, n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "rotate_component")
	fields := internal.NewFields(map[string]string{
		"a": internal.ShortFloat(a),
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "rotate", fields, c...)
}

func (t scadTemplate) Rotate(a, x, y, z float64, n string) scadTemplate {
	return NewRotate(a, x, y, z, n, t)
}
