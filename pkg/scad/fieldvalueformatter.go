package scad

// FieldValueFormatter objects nkow if they are set and what their value should look like in OpenSCAD.
type FieldValueFormatter interface {
	isSet() bool
	fieldValueFormat() string
}
