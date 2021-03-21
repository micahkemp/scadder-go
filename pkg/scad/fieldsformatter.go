package scad

// fieldsFormatter objects know what their fieldsFormat value is.
type fieldsFormatter interface {
	fieldsFormat() string
}
