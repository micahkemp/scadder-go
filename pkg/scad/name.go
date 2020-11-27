package scad

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
)

type name struct {
	defaultName   string
	SpecifiedName string
}

// start with letter, end with letter
// digits and underscores permitted between
const validPattern = "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z0-9]$"

// filename extension
const extension = "scad"

func (n *name) setDefault(s string) {
	n.defaultName = s
}

func (n name) Value() string {
	if n.SpecifiedName != "" {
		return n.SpecifiedName
	}

	return n.defaultName
}

func (n name) Valid() bool {
	return regexp.MustCompile(validPattern).MatchString(n.Value())
}

func (n name) String() string {
	if !n.Valid() {
		log.Fatalf("Invalid name: (%s)\n", n.Value())
	}

	return n.Value()
}

func (n name) DirPath(p string) string {
	return filepath.Join(p, n.String())
}

func (n name) Filename() string {
	return fmt.Sprintf("%s.%s", n, extension)
}

func (n name) FilePath(p string) string {
	return filepath.Join(p, n.Filename())
}
