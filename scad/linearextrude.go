package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewLinearExtrude(name string, h float64, c ...scadTemplate) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "linear_extrude_component")
	fields := internal.NewFields(map[string]string{
		"h": internal.ShortFloat(h),
	})

	return newTemplate(n, internal.TransformationTemplate, "linear_extrude", fields, c...)
}

func (t scadTemplate) LinearExtrude(name string, h float64) scadTemplate {
	return NewLinearExtrude(name, h, t)
}
