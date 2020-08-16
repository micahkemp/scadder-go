package scad

import "github.com/micahkemp/scad/scad/internal"

func NewCube(x, y, z float64, n string) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size": internal.ShortFloatList(x, y, z),
	})

	return newTemplate(name, internal.ShapeTemplate, "cube", fields)
}
