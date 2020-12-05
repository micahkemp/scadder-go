package dice

import "github.com/micahkemp/scad/pkg/scad"

type Die struct {
	Name   string
	Size   float64
	Dimple Dimple
}

var ExampleDie = Die{
	Name: "example_die",
	Size: 50,
	Dimple: Dimple{
		Diameter: 10,
		Depth:    3,
	},
}

func (d Die) SCADWriter() scad.SCADWriter {
	dimples := SideDimples{
		SideLength: d.Size,
		Dimple: Dimple{
			Diameter: d.Dimple.Diameter,
			Depth:    d.Dimple.Depth,
		},
	}

	cube := scad.Cube{
		X:      d.Size,
		Y:      d.Size,
		Z:      d.Size,
		Center: true,
	}

	return scad.DifferenceWithName(
		scad.Name{
			Specified: d.Name,
			Default:   "die",
		}.String(),
		cube,
		dimples,
	).SCADWriter()
}
