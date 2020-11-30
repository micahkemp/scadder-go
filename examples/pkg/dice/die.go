package dice

import "github.com/micahkemp/scad/pkg/scad"

type Die struct {
	Name string
	Size float64
}

var ExampleDie = Die{
	Name: "example_die",
	Size: 10,
}

func (d Die) SCADWriter() scad.SCADWriter {
	return scad.Cube{
		Name: scad.Name{
			Specified: d.Name,
			Default:   "die",
		}.String(),
		X:      d.Size,
		Y:      d.Size,
		Z:      d.Size,
		Center: false,
	}.SCADWriter()
}
