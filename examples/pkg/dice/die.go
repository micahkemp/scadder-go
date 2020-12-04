package dice

import "github.com/micahkemp/scad/pkg/scad"

type Die struct {
	Name           string
	Size           float64
	DimpleDiameter float64
	DimpleDepth    float64
}

var ExampleDie = Die{
	Name:           "example_die",
	Size:           50,
	DimpleDiameter: 10,
	DimpleDepth:    3,
}

func (d Die) SCADWriter() scad.SCADWriter {
	dimples := SideDimples{
		SideLength:     d.Size,
		DimpleDiameter: d.DimpleDiameter,
		DimpleDepth:    d.DimpleDepth,
	}

	cube := scad.Cube{
		X:      d.Size,
		Y:      d.Size,
		Z:      d.Size,
		Center: true,
	}

	return scad.Transformable(cube).SubtractWithName(
		scad.Name{
			Specified: d.Name,
			Default:   "die",
		}.String(),
		dimples,
	).SCADWriter()
}
