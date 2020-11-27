package scad

import "fmt"

type templateWithChildren struct {
	template
	Children []Renderer
}

func (t templateWithChildren) content() string {
	return fmt.Sprintf("%s {%q}", t.template.content(), t.Children)
}

func (t templateWithChildren) Render(path string) {
	for _, child := range t.Children {
		child.Render(t.PathFrom(path))
	}

	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
