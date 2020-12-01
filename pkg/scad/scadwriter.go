package scad

import (
	"bytes"
	"fmt"
	"log"
	tt "text/template"
)

type SCADWriter struct {
	Name
	Function       string
	Fields         fields
	Children       []Module
	templateString string
}

func (t SCADWriter) content() string {
	// our template's name is essentially a throwaway
	templ := tt.Must(tt.New("scad").Parse(t.templateString))
	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, t); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (t SCADWriter) WriteSCAD(path string) {
	for _, child := range t.Children {
		child.SCADWriter().WriteSCAD(t.PathFrom(path))
	}

	fmt.Printf("%s - %s\n", t.Name.SCADFilePath(path), t.content())
}
