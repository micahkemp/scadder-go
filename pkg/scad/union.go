package scad

import (
	"fmt"
)

type Union struct {
	Name     string
	Children []Renderer
	template templateWithChildren
}

func (u Union) Render(path string) {
	name := Name{
		u.Name,
		"union",
	}

	u.template.Name = name
	u.template.Fields = fmt.Sprintf("union()")
	u.template.Children = u.Children
	u.template.render(path)
}
