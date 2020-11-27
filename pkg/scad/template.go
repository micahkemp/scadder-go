package scad

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad/internal"
)

type template struct {
	internal.Name
	Fields string
}

func (t template) content() string {
	return fmt.Sprintf("%s: %s", t.Name, t.Fields)
}

func (t template) render(path string) {
	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
