package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

func NewPolygon(name string, points Points2D) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "polygon_component")
	fields := internal.NewFields(map[string]string{
		"points": points.String(),
	})

	return newTemplate(n, internal.ShapeTemplate, "polygon", fields)
}
