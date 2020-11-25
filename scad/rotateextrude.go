package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

type RotateExtrude struct {
	scadTemplate
}

func NewRotateExtrude(name string, angle float64, c ...DirPathRenderer) RotateExtrude {
	n, _ := internal.FirstNonEmptyName(name, "rotate_extrude_component")
	fields := internal.NewFields(map[string]string{
		"angle": internal.ShortFloat(angle),
	})

	return RotateExtrude{
		newTemplate(n, internal.TransformationTemplate, "rotate_extrude", fields, c...),
	}
}

func (t scadTemplate) RotateExtrude(name string, angle float64) RotateExtrude {
	return NewRotateExtrude(name, angle, t)
}
