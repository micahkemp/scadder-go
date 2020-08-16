package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewSquare(x, y float64, center bool, n string) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "square_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(name, internal.ShapeTemplate, "square", fields)
}
