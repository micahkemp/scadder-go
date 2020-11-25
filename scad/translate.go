package scad

import "github.com/micahkemp/scad/scad/internal"

func NewTranslate(name string, x, y, z float64, c ...DirPathRenderer) ScadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "translate_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(n, internal.TransformationTemplate, "translate", fields, c...)
}

func (t ScadTemplate) Translate(name string, x, y, z float64) ScadTemplate {
	return NewTranslate(name, x, y, z, t)
}
