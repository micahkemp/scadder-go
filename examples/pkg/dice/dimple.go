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
	Depth:    3,
}

func (d Dimple) SCADWriter() scad.SCADWriter {
	return scad.Transformable(scad.Sphere{
		Diameter: d.Diameter,
	}).TranslateWithName(
		scad.Name{
			Specified: d.Name,
			Default:   "dimple",
		}.String(),
		0, 0, d.Depth-d.Diameter/2,
	).SCADWriter()
}
