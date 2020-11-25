package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"strconv"
)

type Square struct {
	scadTemplate
}

func NewSquare(name string, x, y float64, center bool) Square {
	n, _ := internal.FirstNonEmptyName(name, "square_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(x, y),
		"center": strconv.FormatBool(center),
	})

	return Square{
		newTemplate(n, internal.ShapeTemplate, "square", fields),
	}
}
