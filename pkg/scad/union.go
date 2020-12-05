package scad

type Union struct {
	Name     string
	Children []Module
}

func (u Union) SCADWriter() SCADWriter {
	return SCADWriter{
		Name: Name{
			Specified: u.Name,
			Default:   "union_module",
		},
		templateString: transformationTemplate,
		Function:       "union",
		Children:       u.Children,
	}
}

func UnionOf(children ...Module) Union {
	return Union{
		Children: children,
	}
}

func UnionWithName(name string, children ...Module) Union {
	return Union{
		Name:     name,
		Children: children,
	}
}
