package scad

// fieldFormatter is an interface for objects that know if they are set and what their fieldFormat value is.
type fieldFormatter interface {
	isSet() bool
	fieldFormat() string
}
