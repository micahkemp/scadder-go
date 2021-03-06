package scad

import "strings"

// fields is a list of field objects.
type fields []field

// isSet returns true if any of its members are set.
func (f fields) isSet() bool {
	for _, field := range f {
		if field.isSet() {
			return true
		}
	}

	return false
}

// fieldsFormat returns a string representation of fields suitable to use in a function call in OpenSCAD.
func (f fields) fieldsFormat() string {
	fieldStrings := make([]string, 0)

	for _, field := range f {
		if field.isSet() {
			fieldStrings = append(fieldStrings, field.fieldFormat())
		}
	}

	return strings.Join(fieldStrings, ", ")
}
