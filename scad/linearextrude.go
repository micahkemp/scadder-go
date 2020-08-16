package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewLinearExtrude(h float64, n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "linear_extrude_component")
	fields := internal.NewFields(map[string]string{
		"h": internal.ShortFloat(h),
	})

	return newTemplate(name, internal.TransformationTemplate, "linear_extrude", fields, c...)
}

func (t scadTemplate) LinearExtrude(h float64, n string) scadTemplate {
	return NewLinearExtrude(h, n, t)
}
