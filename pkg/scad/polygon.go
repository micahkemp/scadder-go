package scad

type Polygon struct {
	Name   string
	Points XYCoordinates
}

func (p Polygon) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: p.Name,
			Default: "polygon_component",
		},
		templateString: shapeTemplate,
		Function: "polygon",
		Fields: fields{
			"points": p.Points.String(),
		},
	}
}
