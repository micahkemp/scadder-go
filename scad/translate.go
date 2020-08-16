package scad

import "github.com/micahkemp/scad/scad/internal"

func NewTranslate(n string, x, y, z float64, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "translate_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "translate", fields, c...)
}

func (t scadTemplate) Translate(n string, x, y, z float64) scadTemplate {
	return NewTranslate(n, x, y, z, t)
}
