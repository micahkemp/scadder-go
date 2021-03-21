package scad

import (
	"strings"
)

// BoolField is a field that holds a boolean value.
type BoolField struct {
	Name  string
	Value BoolWithExplicitness
}

// SCADString returns the string representation of a BoolField suitable to use as an OpenSCAD parameter.
func (b BoolField) SCADString() string {
	return strings.Join([]string{b.Name, b.Value.SCADString()}, "=")
}
