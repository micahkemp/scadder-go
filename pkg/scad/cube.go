package scad

import (
	"strconv"
)

type Cube struct {
	Name    string
	X, Y, Z float64
	Center  bool
}

func (c Cube) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: c.Name,
			Default:   "cube",
		},
		Function:       "cube",
		templateString: shapeTemplate,
		Fields: fields{
			"size":   shortFloatList(c.X, c.Y, c.Z),
			"center": strconv.FormatBool(c.Center),
		},
	}
}
