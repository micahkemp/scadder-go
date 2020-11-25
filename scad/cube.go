package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCube(name string, x, y, z float64, center bool) ScadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y, z),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(n, internal.ShapeTemplate, "cube", fields)
}
