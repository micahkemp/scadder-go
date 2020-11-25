package scad

import "github.com/micahkemp/scad/scad/internal"

type Mirror struct {
	scadTemplate
}

func NewMirror(name string, x, y, z float64, c ...DirPathRenderer) Mirror {
	n, _ := internal.FirstNonEmptyName(name, "mirror_component")
	fields := internal.NewFields(map[string]string{
		"v": internal.ShortFloatList(x, y, z),
	})

	return Mirror{
		newTemplate(n, internal.TransformationTemplate, "mirror", fields, c...),
	}
}

func (t scadTemplate) Mirror(name string, x, y, z float64) Mirror {
	return NewMirror(name, x, y, z, t)
}
