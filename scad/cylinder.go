package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCylinderByRadius(r, h float64, center bool, n string) scadTemplate {
	name, _ := internal.FirstValidName(n, "cylinder_component")
	fields := internal.NewFields(map[string]string{
		"r": internal.ShortFloat(r),
		"h": internal.ShortFloat(h),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "cylinder", fields)
}

func NewCylinderByDiameter(d, h float64, center bool, n string) scadTemplate {
	return NewCylinderByRadius(d/2, h, center, n)
}
