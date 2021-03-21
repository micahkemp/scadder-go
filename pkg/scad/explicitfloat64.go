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

// SCADString returns a string representation of an explicitBool suitable for use in OpenSCAD.
func (f Float64WithExplicitness) SCADString() string {
	return strconv.FormatFloat(f.float64, 'f', -1, 64)
}
