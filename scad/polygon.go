package scad

import (
	"github.com/micahkemp/scad/scad/internal"
)

type Polygon struct {
	scadTemplate
}

func NewPolygon(name string, points Points2D) Polygon {
	n, _ := internal.FirstNonEmptyName(name, "polygon_component")
	fields := internal.NewFields(map[string]string{
		"points": points.String(),
	})

	return Polygon{
		newTemplate(n, internal.ShapeTemplate, "polygon", fields),
	}
}
