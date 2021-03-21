package scad

import "fmt"

// field represents a key=value parameter to pass to OpenSCAD.
type field struct {
	name string
	isSetStringer
}

// SCADString returns a string representation of a field suitable to use in OpenSCAD.
func (f field) SCADString() string {
	return fmt.Sprintf("%s=%s", f.name, f.isSetStringer.SCADString())
}
