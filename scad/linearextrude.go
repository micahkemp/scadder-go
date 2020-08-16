package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewLinearExtrude(n string, h float64, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "linear_extrude_component")
	fields := internal.NewFields(map[string]string{
		"h": internal.ShortFloat(h),
	})

	return newTemplate(name, internal.TransformationTemplate, "linear_extrude", fields, c...)
}

func (t scadTemplate) LinearExtrude(n string, h float64) scadTemplate {
	return NewLinearExtrude(n, h, t)
}
