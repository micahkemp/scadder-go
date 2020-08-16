package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCircleByRadius(n string, r float64, center bool) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "circle_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "circle", fields)
}

func NewCircleByDiameter(n string, d float64, center bool) scadTemplate {
	return NewCircleByRadius(n, d/2, center)
}
