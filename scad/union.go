package scad

import "github.com/micahkemp/scad/scad/internal"

type Union struct {
	scadTemplate
}

func NewUnion(name string, c ...DirPathRenderer) Union {
	n, _ := internal.FirstNonEmptyName(name, "union_component")

	return Union{
		newTemplate(n, internal.TransformationTemplate, "union", internal.Fields{}, c...),
	}
}

func (t scadTemplate) Add(name string, c ...DirPathRenderer) Union {
	allChildren := append([]DirPathRenderer{t}, c...)
	return NewUnion(name, allChildren...)
}
