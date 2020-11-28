package scad

import (
	"strconv"
)

type Cube struct {
	Name     string
	X, Y, Z  float64
	Center   bool
	template template
}

func (c Cube) Render(path string) {
	c.template.Name = Name{
		Specified: c.Name,
		Default:   "cube",
	}

	c.template.fields = fields{
		"size": shortFloatList(c.X, c.Y, c.Z),
		"center": strconv.FormatBool(c.Center),
	}
	c.template.render(path)
}
