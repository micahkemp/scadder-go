package scad

type Translate struct {
	Name     string
	X, Y, Z  float64
	Children []Renderer
	template templateWithChildren
}

func (t Translate) Render(path string) {
	name := Name{
		Specified: t.Name,
		Default:   "translate",
	}

	t.template.Name = name
	t.template.fields = fields{
		"v": shortFloatList(t.X, t.Y, t.Z),
	}
	t.template.Children = t.Children
	t.template.render(path)
}

func (r renderable) Translate(x, y, z float64) Translate {
	return Translate{
		X: x,
		Y: y,
		Z: z,
		Children: []Renderer{
			r.renderer,
		},
	}
}

func (r renderable) TranslateWithName(n string, x, y, z float64) Translate {
	t := r.Translate(x, y, z)
	t.Name = n

	return t
}
