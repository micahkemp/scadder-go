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

func (d Die) Render(path string) {
	name := scad.Name{
		Specified: d.Name,
		Default:   "die",
	}

	cube := scad.Cube{
		Name:   name.String(),
		X:      d.Size,
		Y:      d.Size,
		Z:      d.Size,
		Center: false,
	}

	cube.Render(path)
}
