package scad

import "github.com/micahkemp/scad/scad/internal"

type Translate struct {
	scadTemplate
}

func NewTranslate(name string, x, y, z float64, c ...DirPathRenderer) Translate {
	n, _ := internal.FirstNonEmptyName(name, "translate_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return Translate{
		newTemplate(n, internal.TransformationTemplate, "translate", fields, c...),
	}
}

func (t scadTemplate) Translate(name string, x, y, z float64) Translate {
	return NewTranslate(name, x, y, z, t)
}
