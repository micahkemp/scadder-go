package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

type Cube struct {
	scadTemplate
}

func NewCube(name string, x, y, z float64, center bool) Cube {
	n, _ := internal.FirstNonEmptyName(name, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y, z),
		"center": strconv.FormatBool(center),
	})

	return Cube{
		newTemplate(n, internal.ShapeTemplate, "cube", fields),
	}
}
