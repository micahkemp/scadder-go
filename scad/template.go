package scad

import (
	"bytes"
	"fmt"
	"github.com/micahkemp/scad/scad/internal"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type scadTemplate struct {
	Name     internal.Name
	template string
	CallName string
	Fields   internal.Fields
	Children []scadTemplate
}

func NewTemplate(name internal.Name, template string, callName string, fields internal.Fields, children ...scadTemplate) scadTemplate {
	return scadTemplate{
		name,
		template,
		callName,
		fields,
		children,
	}
}

func (t scadTemplate) String() string {
	// our template's name is essentially a throwaway
	templ := template.Must(template.New("scad").Parse(t.template))
	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, t); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (t scadTemplate) rendered(path string) bool {
	contents, err := ioutil.ReadFile(t.Name.FilePath(path))

	if err != nil {
		// any error assumes not rendered
		return false
	}

	return string(contents) == t.String()
}

func (t scadTemplate) childPath(child scadTemplate, path string) string {
	return filepath.Join(path, child.Name.String())
}

func (t scadTemplate) childRendered(child scadTemplate, path string) bool {
	return child.rendered(t.childPath(child, path))
}

func (t scadTemplate) Render(path string) (ok bool) {
	if err := os.MkdirAll(path, 0755); err != nil {
		// TODO this needs to be handled properly
		fmt.Printf("Unable to create directory: %s: %s", path, err)
		return false
	}

	// handle children first
	for _, c := range t.Children {
		if cRenderOk := c.Render(t.childPath(c, path)); cRenderOk == false {
			return false
		}
	}

	if t.rendered(path) {
		return true
	}

	// OpenFile used specifically to enforce not overwriting an existing file, but instead failing
	f, err := os.OpenFile(t.Name.FilePath(path), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	// TODO this needs to be handled properly
	if err != nil {
		fmt.Printf("(scadTemplate) Render(%q): %s\n", path, err)
		return false
	}

	defer f.Close()

	if _, err = f.WriteString(t.String()); err != nil {
		fmt.Printf("(scadTemplate) Render(%q): %s\n", path, err)
		return false
	}

	return true
}
