package scad

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad/internal"
)

type Cube struct {
	Name     string
	X, Y, Z  float64
	Center   bool
	template template
}

func (c Cube) Render(path string) {
	c.template.Name = internal.Name{
		Specified: c.Name,
		Default:   "cube",
	}

	c.template.Fields = fmt.Sprintf("cube([%f, %f, %f, %v])", c.X, c.Y, c.Z, c.Center)
	c.template.Render(path)
}
