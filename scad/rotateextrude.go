package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewRotateExtrude(angle float64, n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "rotate_extrude_component")
	fields := internal.NewFields(map[string]string{
		"angle": internal.ShortFloat(angle),
	})

	return newTemplate(name, internal.TransformationTemplate, "rotate_extrude", fields, c...)
}

func (t scadTemplate) RotateExtrude(angle float64, n string) scadTemplate {
	return NewRotateExtrude(angle, n, t)
}
