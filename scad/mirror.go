package scad

import "github.com/micahkemp/scad/scad/internal"

func NewMirror(name string, x, y, z float64, c ...scadTemplate) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "mirror_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(n, internal.TransformationTemplate, "mirror", fields, c...)
}

func (t scadTemplate) Mirror(name string, x, y, z float64) scadTemplate {
	return NewMirror(name, x, y, z, t)
}
