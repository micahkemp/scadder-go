package scad

import (
	"fmt"
	"log"
	"path/filepath"
)

type Name struct {
	Specified string
	Default   string
}

func (n Name) value() (value string, ok bool) {
	value = n.Default
	ok = true

	if n.Specified != "" {
		value = n.Specified
	}

	if value == "" {
		ok = false
	}

	return
}

func (n Name) String() string {
	value, ok := n.value()
	if !ok {
		log.Fatal("Attempted to stringify empty Name")
	}

	return value
}

func (n Name) PathFrom(path string) string {
	return filepath.Join(path, n.String())
}

func (n Name) SCADFileName() string {
	return fmt.Sprintf("%s.scad", n.String())
}

func (n Name) SCADFilePath(path string) string {
	return filepath.Join(path, n.SCADFileName())
}
