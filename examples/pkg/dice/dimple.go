package dice

import "github.com/micahkemp/scad/pkg/scad"

type Dimple struct {
	Name     string
	Diameter float64
	Depth    float64
}

var ExampleDimple = Dimple{
	Name:     "example_dimple",
	Diameter: 10,
	Depth:    8,
}

func (d Dimple) SCADWriter() scad.SCADWriter {
	sphere := scad.Transformable(scad.Sphere{
		Diameter: d.Diameter,
	}).Translate(0, 0, d.Depth-d.Diameter)

	subtractCube := scad.Transformable(scad.Cube{
		X:      d.Diameter,
		Y:      d.Diameter,
		Z:      d.Diameter,
		Center: true,
	}).TranslateWithName("subtract_cube", 0, 0, -d.Diameter/2)

	return scad.Difference{
		Name: scad.Name{
			Specified: d.Name,
			Default:   "dimple",
		}.String(),
		Children: []scad.Module{
			sphere,
			subtractCube,
		},
	}.SCADWriter()
}
