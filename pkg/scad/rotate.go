package scad

type Rotate struct {
	Name     string
	Angle    float64
	X, Y, Z  float64
	Children []Module
}

func (r Rotate) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: r.Name,
			Default:   "rotate_component",
		},
		templateString: transformationTemplate,
		Function:       "rotate",
		Fields: fields{
			"a": shortFloat(r.Angle),
			"v": shortFloatList(r.X, r.Y, r.Z),
		},
		Children: r.Children,
	}
}

func RotationOf(a, x, y, z float64, children ...Module) Rotate {
	return Rotate{
		Angle:    a,
		X:        x,
		Y:        y,
		Z:        z,
		Children: children,
	}
}

func RotationWithName(name string, a, x, y, z float64, children ...Module) Rotate {
	return Rotate{
		Name:     name,
		Angle:    a,
		X:        x,
		Y:        y,
		Z:        z,
		Children: children,
	}
}
