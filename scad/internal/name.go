package internal

import (
	"fmt"
	"path/filepath"
	"regexp"
)

type Name string

// start with letter, end with letter
// digits and underscores permitted between
const validPattern = "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z0-9]$"

// filename extension
const extension = "scad"

func valid(s string) bool {
	return regexp.MustCompile(validPattern).MatchString(s)
}

func NewName(s string) (n Name, ok bool) {
	if valid(s) {
		n = Name(s)
		ok = true
	}

	return
}

// returns
//   name: first non-empty name passed
//   ok:   the name's validity
func FirstNonEmptyName(names ...string) (name string, ok bool) {
	for _, name = range names {
		if name != "" {
			_, ok = NewName(name)
			return
		}
	}

	// clear name before returning with ok = false
	name = ""
	return
}

func (n Name) String() string {
	return string(n)
}

func (n Name) Filename() string {
	return fmt.Sprintf("%s.%s", n, extension)
}

func (n Name) FilePath(p string) string {
	return filepath.Join(p, n.Filename())
}
