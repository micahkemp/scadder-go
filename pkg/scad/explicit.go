package scad

// explicit is used to determine if a value was set explicitly, even if it would otherwise evaluate to false.
type explicit bool

// isSet returns true if explicit is set to true. When explicit is an anonymous member of a type, it can be called
// directly on an object of that type to determine if the value was set explicitly.
func (e explicit) isSet() bool {
	return bool(e)
}
