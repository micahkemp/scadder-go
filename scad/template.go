package scad

import (
	"bytes"
	"log"
	"text/template"
)

type Template struct {
	Name     Name
	template string
	CallName string
	Fields   Fields
	Children []Template
}

func NewTemplate(name Name, template string, callName string, fields Fields, children ...Template) Template {
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
