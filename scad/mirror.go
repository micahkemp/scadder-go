package scad

import "github.com/micahkemp/scad/scad/internal"

func NewMirror(x, y, z float64, n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "mirror_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "mirror", fields, c...)
}

func (t scadTemplate) Mirror(x, y, z float64, n string) scadTemplate {
	return NewMirror(x, y, z, n, t)
}
