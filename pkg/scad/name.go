package scad

import (
	"fmt"
	"log"
	"path/filepath"
)

// Name is a name of a SCAD component. To be valid it must have at least one of specified or default set.
type Name struct {
	specifiedValue string
	defaultValue   string
}

// value returns a Name's specifiedValue if present, the defaultValue otherwise. It also returns a boolean ok to indicate
// if the returned value is valid.
func (n Name) value() (value string, ok bool) {
	value = n.defaultValue
	ok = true

	if n.specifiedValue != "" {
		value = n.specifiedValue
	}

	if value == "" {
		ok = false
	}

	return
}

// String returns a name's value, and logs a fatal error if it was attempted against an invalid Name.
func (n Name) String() string {
	value, ok := n.value()
	if !ok {
		log.Fatal("Attempted to stringify empty Name")
	}

	return value
}

// PathFrom returns the full path for a Name from a passed path.
func (n Name) PathFrom(path string) string {
	return filepath.Join(path, n.String())
}

// FileName returns a Name's filename.
func (n Name) FileName() string {
	return fmt.Sprintf("%s.scad", n.String())
}

// FilePath returns a Name's path (including filename) from a passed path.
func (n Name) FilePath(path string) string {
	return filepath.Join(path, n.FileName())
}
