package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

type Circle struct {
	scadTemplate
}

func NewCircleByRadius(name string, r float64, center bool) Circle {
	n, _ := internal.FirstNonEmptyName(name, "circle_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"center": strconv.FormatBool(center),
	})

	return Circle{
		newTemplate(n, internal.ShapeTemplate, "circle", fields),
	}
}

func NewCircleByDiameter(name string, d float64, center bool) Circle {
	return NewCircleByRadius(name, d/2, center)
}
