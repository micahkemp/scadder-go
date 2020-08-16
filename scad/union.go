package scad

import "github.com/micahkemp/scad/scad/internal"

func NewUnion(n string, c ...scadTemplate) scadTemplate {
	name, _ := internal.FirstNonEmptyName(n, "difference_component")

	return newTemplate(name, internal.TransformationTemplate, "union", internal.Fields{}, c...)
}

func (t scadTemplate) Add(n string, c ...scadTemplate) scadTemplate {
	allChildren := append([]scadTemplate{t}, c...)
	return NewUnion(n, allChildren...)
}
