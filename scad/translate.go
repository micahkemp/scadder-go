package scad

func NewTranslate(x, y, z float64, n string, c ...Template) Template {
	name, _ := FirstValidName(n, "translate_module")
	fields := NewFields(map[string]string{
		"v": ShortFloatList(x, y, z),
	})

	return NewTemplate(name, transformationTemplate, "translate", fields, c...)
}
