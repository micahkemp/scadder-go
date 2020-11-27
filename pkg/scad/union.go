package scad

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad/internal"
)

type Union struct {
	Name     string
	Children []Renderer
	template TemplateWithChildren
}

func (u Union) Render(path string) {
	name := internal.Name{
		u.Name,
		"union",
	}

	u.template.Name = name
	u.template.Fields = fmt.Sprintf("union()")
	u.template.Children = u.Children
	u.template.Render(path)
}
