package scad

import "github.com/micahkemp/scad/scad/internal"

func NewMirror(n string, x, y, z float64, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "mirror_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "mirror", fields, c...)
}

func (t scadTemplate) Mirror(n string, x, y, z float64) scadTemplate {
	return NewMirror(n, x, y, z, t)
}
