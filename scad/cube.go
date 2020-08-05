package scad

import "github.com/micahkemp/scad/scad/internal"

func NewCube(x, y, z float64, n string) Template {
	name, _ := internal.FirstValidName(n, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size": internal.ShortFloatList(x, y, z),
	})

	return NewTemplate(name, shapeTemplate, "cube", fields)
}
