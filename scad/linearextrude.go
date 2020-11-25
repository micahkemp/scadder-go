package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewLinearExtrude(name string, h float64, c ...DirPathRenderer) ScadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "linear_extrude_component")
	fields := internal.NewFields(map[string]string{
		"h": internal.ShortFloat(h),
	})

	return newTemplate(n, internal.TransformationTemplate, "linear_extrude", fields, c...)
}

func (t ScadTemplate) LinearExtrude(name string, h float64) ScadTemplate {
	return NewLinearExtrude(name, h, t)
}
