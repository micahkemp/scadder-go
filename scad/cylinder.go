package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCylinderByRadius(name string, r, h float64, center bool) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "cylinder_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"h":      internal.ShortFloat(h),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(n, internal.ShapeTemplate, "cylinder", fields)
}

func NewCylinderByDiameter(name string, d, h float64, center bool) scadTemplate {
	return NewCylinderByRadius(name, d/2, h, center)
}
