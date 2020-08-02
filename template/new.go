package template

import (
	"bytes"
	"github.com/micahkemp/scad/fields"
	"github.com/micahkemp/scad/name"
	"log"
	"text/template"
)

type Template struct {
	Name     name.Name
	template string
	CallName string
	Fields   fields.Fields
}

func New(n name.Name, s string, c string, f fields.Fields) Template {
	return Template{
		n,
		s,
		c,
		f,
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
