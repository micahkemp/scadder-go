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
}

func NewTemplate(n Name, s string, c string, f Fields) Template {
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
