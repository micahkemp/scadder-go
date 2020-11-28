package scad

type Union struct {
	Name     string
	Children []Renderer
	template templateWithChildren
}

func (u Union) Render(path string) {
	name := Name{
		Specified: u.Name,
		Default:   "union",
	}

	u.template.Name = name
	u.template.Children = u.Children
	u.template.render(path)
}

func (r renderable) Add(c ...Renderer) Union {
	allChildren := append([]Renderer{r.renderer}, c...)

	return Union{
		Children: allChildren,
	}
}

func (r renderable) AddWithName(n string, c ...Renderer) Union {
	u := r.Add(c...)

	u.Name = n

	return u
}
