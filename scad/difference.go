package scad

import "github.com/micahkemp/scad/scad/internal"

type Difference struct {
	scadTemplate
}

func NewDifference(name string, c ...DirPathRenderer) Difference {
	n, _ := internal.FirstNonEmptyName(name, "difference_component")

	return Difference{
		newTemplate(n, internal.TransformationTemplate, "difference", internal.Fields{}, c...),
	}
}

// Subtract only takes one scadTemplate argument, because of oddness of difference() when passed more than two children
func (t scadTemplate) Subtract(name string, c DirPathRenderer) Difference {
	return NewDifference(name, t, c)
}
