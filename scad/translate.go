package scad

import "github.com/micahkemp/scad/scad/internal"

func NewTranslate(x, y, z float64, n string, c ...Template) Template {
	name, _ := internal.FirstValidName(n, "translate_module")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return NewTemplate(name, transformationTemplate, "translate", fields, c...)
}
