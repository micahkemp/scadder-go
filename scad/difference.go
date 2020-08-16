package scad

import "github.com/micahkemp/scad/scad/internal"

func NewDifference(n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "difference_component")

	return newTemplate(name, internal.TransformationTemplate, "difference", internal.Fields{}, c...)
}

func (t scadTemplate) Subtract(n string, c scadTemplate) scadTemplate {
	return NewDifference(n, t, c)
}
