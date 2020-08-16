package scad

import "github.com/micahkemp/scad/scad/internal"

func NewSquare(x, y float64, n string) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "square_component")
	fields := internal.NewFields(map[string]string{
		"size": internal.ShortFloatList(x, y),
	})

	return newTemplate(name, internal.ShapeTemplate, "square", fields)
}
