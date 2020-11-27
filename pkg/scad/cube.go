package scad

import (
	"github.com/micahkemp/scad/pkg/scad/internal"
	"strconv"
)

type Cube struct {
	Name    string
	X, Y, Z float64
	Center  bool
	ScadTemplate
}

var CubeExample = Cube{
	Name:   "cube_example",
	X:      10,
	Y:      20,
	Z:      30,
	Center: true,
}

func NewCube(n string, x, y, z float64, center bool) Cube {
	cube := Cube{
		Name:   n,
		X:      x,
		Y:      y,
		Z:      z,
		Center: center,
	}
	cube.Initialize()

	return cube
}

func (c *Cube) Initialize() {
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(c.X, c.Y, c.Z),
		"center": strconv.FormatBool(c.Center),
	})

	cubeName := name{
		SpecifiedName: c.Name,
		defaultName:   "cube_component",
	}
	c.ScadTemplate = NewTemplate(cubeName, internal.ShapeTemplate, "cube", fields)
}
