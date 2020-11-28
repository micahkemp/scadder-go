package scad

import (
	"fmt"
)

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
	t.template.Fields = fmt.Sprintf("[%f, %f, %f]", t.X, t.Y, t.Z)
	t.template.Children = t.Children
	t.template.render(path)
}
