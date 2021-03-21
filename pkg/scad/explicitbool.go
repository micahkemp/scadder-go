package scad

import "strconv"

// BoolWithExplicitness is a boolean that can be set explicitly.
type BoolWithExplicitness struct {
	bool
	explicit
}

// ExplicitBool returns a BoolWithExplicitness that has been set explicitly.
func ExplicitBool(value bool) BoolWithExplicitness {
	return BoolWithExplicitness{
		value,
		true,
	}
}

// fieldValueFormat returns a string representation of a BoolWithExplicitness suitable for use in OpenSCAD.
func (e BoolWithExplicitness) fieldValueFormat() string {
	return strconv.FormatBool(e.bool)
}
