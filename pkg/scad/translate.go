package scad

type Translate struct {
	Name     string
	X, Y, Z  float64
	Children []Module
}

func (t Translate) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: t.Name,
			Default:   "translate_module",
		},
		templateString: transformationTemplate,
		Function:       "translate",
		Fields: fields{
			"v": shortFloatList(t.X, t.Y, t.Z),
		},
		Children: t.Children,
	}
}

func TranslationOf(x, y, z float64, children ...Module) Translate {
	return Translate{
		X:        x,
		Y:        y,
		Z:        z,
		Children: children,
	}
}

func TranslationWithName(name string, x, y, z float64, children ...Module) Translate {
	return Translate{
		Name:     name,
		X:        x,
		Y:        y,
		Z:        z,
		Children: children,
	}
}
