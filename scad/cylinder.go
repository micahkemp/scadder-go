package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCylinderByRadius(n string, r, h float64, center bool) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "cylinder_component")
	fields := internal.NewFields(map[string]string{
		"r":      internal.ShortFloat(r),
		"h":      internal.ShortFloat(h),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "cylinder", fields)
}

func NewCylinderByDiameter(n string, d, h float64, center bool) scadTemplate {
	return NewCylinderByRadius(n, d/2, h, center)
}
