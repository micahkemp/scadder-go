package scad

import "github.com/micahkemp/scad/scad/internal"

func NewUnion(name string, c ...DirPathRenderer) scadTemplate {
	n, _ := internal.FirstNonEmptyName(name, "union_component")

	return newTemplate(n, internal.TransformationTemplate, "union", internal.Fields{}, c...)
}

func (t scadTemplate) Add(name string, c ...DirPathRenderer) scadTemplate {
	allChildren := append([]DirPathRenderer{t}, c...)
	return NewUnion(name, allChildren...)
}
