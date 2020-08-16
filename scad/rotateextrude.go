package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewRotateExtrude(name string, angle float64, c ...scadTemplate) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "rotate_extrude_component")
	fields := internal.NewFields(map[string]string{
		"angle": internal.ShortFloat(angle),
	})

	return newTemplate(n, internal.TransformationTemplate, "rotate_extrude", fields, c...)
}

func (t scadTemplate) RotateExtrude(name string, angle float64) scadTemplate {
	return NewRotateExtrude(name, angle, t)
}
