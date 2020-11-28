package scad

import (
	"fmt"
)

type template struct {
	Name
	fields
}

func (t template) content() string {
	return fmt.Sprintf("%s: %q", t.Name, t.fields)
}

func (t template) render(path string) {
	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
