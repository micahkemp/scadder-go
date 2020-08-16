package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCircleByRadius(r float64, center bool, n string) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "circle_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "circle", fields)
}

func NewCircleByDiameter(d float64, center bool, n string) scadTemplate {
	return NewCircleByRadius(d/2, center, n)
}
