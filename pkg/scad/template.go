package scad

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad/internal"
)

type template struct {
	internal.Name
	Fields string
}

type templateWithChildren struct {
	template
	Children []Renderer
}

func (t template) content() string {
	return fmt.Sprintf("%s: %s", t.Name, t.Fields)
}

func (t templateWithChildren) content() string {
	return fmt.Sprintf("%s {%q}", t.template.content(), t.Children)
}

func (t template) Render(path string) {
	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}

func (t templateWithChildren) Render(path string) {
	for _, child := range t.Children {
		child.Render(t.PathFrom(path))
	}

	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
