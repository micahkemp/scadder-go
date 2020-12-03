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

func (t transformable) Add(c ...Module) Union {
	allChildren := append([]Module{t.Module}, c...)

	return Union{
		Children: allChildren,
	}
}

func (t transformable) AddWithName(n string, c ...Module) Union {
	u := t.Add(c...)

	u.Name = n

	return u
}
