package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewCube(x, y, z float64, center bool, n string) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y, z),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "cube", fields)
}
