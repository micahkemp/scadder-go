package scad

import "strconv"

// Float64WithExplicitness is a float64 that can be set explicitly.
type Float64WithExplicitness struct {
	float64
	explicit
}

// ExplicitFloat64 returns a Float64WithExplicitness that has been set explicitly.
func ExplicitFloat64(value float64) Float64WithExplicitness {
	return Float64WithExplicitness{
		value,
		true,
	}
}

// fieldValueFormat returns a string representation of a Float64WithExplicitness suitable for use in OpenSCAD.
func (f Float64WithExplicitness) fieldValueFormat() string {
	return strconv.FormatFloat(f.float64, 'f', -1, 64)
}
