package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

func NewSquare(name string, x, y float64, center bool) ScadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "square_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y),
		"center": strconv.FormatBool(center),
	})

	return newTemplate(n, internal.ShapeTemplate, "square", fields)
}
