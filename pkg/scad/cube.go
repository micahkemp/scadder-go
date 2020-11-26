package scad

import (
	"github.com/micahkemp/scad/pkg/scad/internal"
	"strconv"
)

type Cube struct {
	Name    string
	X, Y, Z float64
	Center  bool
	internal.ScadTemplate
}

var CubeExample = Cube{
	Name:   "cube_example",
	X:      10,
	Y:      20,
	Z:      30,
	Center: true,
}

func NewCube(name string, x, y, z float64, center bool) Cube {
	cube := Cube{
		Name:   name,
		X:      x,
		Y:      y,
		Z:      z,
		Center: center,
	}
	cube.Initialize()

	return cube
}

func (c *Cube) Initialize() {
	n, _ := internal.FirstNonEmptyName(c.Name, "cube_component")
	fields := internal.NewFields(map[string]string{
		"size":   internal.ShortFloatList(c.X, c.Y, c.Z),
		"center": strconv.FormatBool(c.Center),
	})

	c.ScadTemplate = internal.NewTemplate(n, internal.ShapeTemplate, "cube", fields)
}
