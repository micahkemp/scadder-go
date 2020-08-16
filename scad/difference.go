package scad

import "github.com/micahkemp/scad/scad/internal"

func NewDifference(name string, c ...scadTemplate) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "difference_component")

	return newTemplate(n, internal.TransformationTemplate, "difference", internal.Fields{}, c...)
}

// Subtract only takes one scadTemplate argument, because of oddness of difference() when passed more than two children
func (t scadTemplate) Subtract(name string, c scadTemplate) scadTemplate {
	return NewDifference(name, t, c)
}
