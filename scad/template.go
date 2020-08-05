package scad

import (
	"bytes"
	"github.com/micahkemp/scad/scad/internal"
	"log"
	"text/template"
)

type Template struct {
	Name     internal.Name
	template string
	CallName string
	Fields   internal.Fields
	Children []Template
}

func NewTemplate(name internal.Name, template string, callName string, fields internal.Fields, children ...Template) Template {
	return Template{
		name,
		template,
		callName,
		fields,
		children,
	}
}

func (t Template) String() string {
	// our template's name is essentially a throwaway
	templ := template.Must(template.New("scad").Parse(t.template))
	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, t); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
