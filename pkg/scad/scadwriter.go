package scad

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
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

func (t SCADWriter) rendered(path string) bool {
	contents, err := ioutil.ReadFile(t.Name.SCADFilePath(path))

	if err != nil {
		// any error assumes not rendered
		return false
	}


	return string(contents) == t.content()
}

func (t SCADWriter) WriteSCAD(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("Unable to create directory: %s: %s", path, err)
	}

	// handle children first
	for _, c := range t.Children {
		c.SCADWriter().WriteSCAD(c.SCADWriter().Name.PathFrom(path))
	}

	if t.rendered(path) {
		return
	}

	// OpenFile used specifically to enforce not overwriting an existing file, but instead failing
	f, err := os.OpenFile(t.Name.SCADFilePath(path), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	// TODO this needs to be handled properly
	if err != nil {
		log.Fatalf("(scadTemplate) Render(%q): %s\n", path, err)
	}

	defer f.Close()

	if _, err = f.WriteString(t.content()); err != nil {
		log.Fatalf("(scadTemplate) Render(%q): %s\n", path, err)
	}
}
