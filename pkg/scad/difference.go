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

func DifferenceOf(children ...Module) Difference {
	return Difference{
		Children: children,
	}
}

func DifferenceWithName(name string, children ...Module) Difference {
	return Difference{
		Name:     name,
		Children: children,
	}
}
