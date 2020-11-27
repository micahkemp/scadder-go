package scad

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad/internal"
)

type Template struct {
	internal.Name
	Fields string
}

type TemplateWithChildren struct {
	Template
	Children []Renderer
}

func (t Template) content() string {
	return fmt.Sprintf("%s: %s", t.Name, t.Fields)
}

func (t TemplateWithChildren) content() string {
	return fmt.Sprintf("%s {%q}", t.Template.content(), t.Children)
}

func (t Template) Render(path string) {
	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}

func (t TemplateWithChildren) Render(path string) {
	for _, child := range t.Children {
		child.Render(t.PathFrom(path))
	}

	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
