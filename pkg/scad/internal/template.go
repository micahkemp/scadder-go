package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type DirPathRenderer interface {
	DirPath(p string) string
	Render(path string) (ok bool)
}

type ScadTemplate struct {
	Name
	template string
	callName string
	fields   Fields
	children []DirPathRenderer
}

func NewTemplate(name string, template string, callName string, fields Fields, children ...DirPathRenderer) ScadTemplate {
	// TODO - need to handle potential error here
	newName, _ := NewName(name)

	return ScadTemplate{
		newName,
		template,
		callName,
		fields,
		children,
	}
}

func (t ScadTemplate) CallName() string {
	return t.callName
}

func (t ScadTemplate) Fields() Fields {
	return t.fields
}

func (t ScadTemplate) Children() []DirPathRenderer {
	return t.children
}

func (t ScadTemplate) String() string {
	// our template's name is essentially a throwaway
	templ := template.Must(template.New("scad").Parse(t.template))
	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, t); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (t ScadTemplate) rendered(path string) bool {
	contents, err := ioutil.ReadFile(t.Name.FilePath(path))

	if err != nil {
		// any error assumes not rendered
		return false
	}

	return string(contents) == t.String()
}

// childRendered is currently only used for testing
// it should be considered for removal
func (t ScadTemplate) childRendered(child ScadTemplate, path string) bool {
	return child.rendered(child.DirPath(path))
}

func (t ScadTemplate) Render(path string) (ok bool) {
	if err := os.MkdirAll(path, 0755); err != nil {
		// TODO this needs to be handled properly
		fmt.Printf("Unable to create directory: %s: %s", path, err)
		return false
	}

	// handle children first
	for _, c := range t.children {
		if cRenderOk := c.Render(c.DirPath(path)); cRenderOk == false {
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
		fmt.Printf("(ScadTemplate) Render(%q): %s\n", path, err)
		return false
	}

	defer f.Close()

	if _, err = f.WriteString(t.String()); err != nil {
		fmt.Printf("(ScadTemplate) Render(%q): %s\n", path, err)
		return false
	}

	return true
}
