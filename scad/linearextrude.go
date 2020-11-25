package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

type LinearExtrude struct {
	scadTemplate
}

func NewLinearExtrude(name string, h float64, c ...DirPathRenderer) LinearExtrude {
	n, _ := internal.FirstNonEmptyName(name, "linear_extrude_component")
	fields := internal.NewFields(map[string]string{
		"h": internal.ShortFloat(h),
	})

	return LinearExtrude{
		newTemplate(n, internal.TransformationTemplate, "linear_extrude", fields, c...),
	}
}

func (t scadTemplate) LinearExtrude(name string, h float64) LinearExtrude {
	return NewLinearExtrude(name, h, t)
}
