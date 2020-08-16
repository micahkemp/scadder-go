package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewRotateExtrude(n string, angle float64, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "rotate_extrude_component")
	fields := internal.NewFields(map[string]string{
		"angle": internal.ShortFloat(angle),
	})

	return newTemplate(name, internal.TransformationTemplate, "rotate_extrude", fields, c...)
}

func (t scadTemplate) RotateExtrude(n string, angle float64) scadTemplate {
	return NewRotateExtrude(n, angle, t)
}
