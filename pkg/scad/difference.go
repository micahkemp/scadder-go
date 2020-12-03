package scad

type Difference struct {
	Name     string
	Children []Module
}

func (d Difference) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: d.Name,
			Default:   "difference_module",
		},
		templateString: transformationTemplate,
		Function:       "difference",
		Children:       d.Children,
	}
}

func (t transformable) Subtract(c ...Module) Difference {
	allChildren := append([]Module{t.Module}, c...)

	return Difference{
		Children: allChildren,
	}
}

func (t transformable) SubtractWithName(n string, c ...Module) Difference {
	d := t.Subtract(c...)
	d.Name = n
	return d
}
