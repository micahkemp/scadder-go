package scad

import "github.com/micahkemp/scad/scad/internal"

func NewTranslate(x, y, z float64, n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "translate_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.TransformationTemplate, "translate", fields, c...)
}

func (t scadTemplate) Translate(x, y, z float64, n string) scadTemplate {
	return NewTranslate(x, y, z, n, t)
}
