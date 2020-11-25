package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCircleByRadius(name string, r float64, center bool) ScadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "circle_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(n, internal.ShapeTemplate, "circle", fields)
}

func NewCircleByDiameter(name string, d float64, center bool) ScadTemplate {
	return NewCircleByRadius(name, d/2, center)
}
