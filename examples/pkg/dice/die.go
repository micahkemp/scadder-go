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
	cube := scad.Cube{
		Name:   d.Name,
		X:      d.Size,
		Y:      d.Size,
		Z:      d.Size,
		Center: false,
	}

	cube.Render(path)
}
