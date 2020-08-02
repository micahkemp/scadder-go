package scad

func NewCube(x, y, z float64, n string) Template {
	name, _ := FirstValidName(n, "cube")
	fields := NewFields(map[string]string{
		"size": ShortFloatList(x, y, z),
	})

	return NewTemplate(name, shapeTemplate, "cube", fields)
}
