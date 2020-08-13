package scad

import (
	"bytes"
	"github.com/micahkemp/scad/scad/internal"
	"log"
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

func (t scadTemplate) childPath(child scadTemplate, path string) string {
	return filepath.Join(path, child.Name.String())
}
