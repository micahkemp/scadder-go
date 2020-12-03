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

func (r transformable) Translate(x, y, z float64) Translate {
	return Translate{
		X:        x,
		Y:        y,
		Z:        z,
		Children: []Module{r.Module},
	}
}

func (r transformable) TranslateWithName(n string, x, y, z float64) Translate {
	t := r.Translate(x, y, z)
	t.Name = n

	return t
}
