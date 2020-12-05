package scad

import "strconv"

type LinearExtrude struct {
	Name     string
	Height   float64
	Center   bool
	Children []Module
}

func (l LinearExtrude) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: l.Name,
			Default: "linear_extrude_component",
		},
		templateString: transformationTemplate,
		Function: "linear_extrude",
		Fields: fields{
			"height": shortFloat(l.Height),
			"center": strconv.FormatBool(l.Center),
		},
		Children: l.Children,
	}
}
