package scad

import "github.com/micahkemp/scad/scad/internal"

func NewUnion(name string, c ...scadTemplate) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "union_component")

	return newTemplate(n, internal.TransformationTemplate, "union", internal.Fields{}, c...)
}

func (t scadTemplate) Add(name string, c ...scadTemplate) scadTemplate {
	allChildren := append([]scadTemplate{t}, c...)
	return NewUnion(name, allChildren...)
}
