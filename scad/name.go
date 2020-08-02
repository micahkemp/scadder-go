package scad

import (
	"fmt"
	"regexp"
)

type Name string

// start with letter, end with letter
// digits and underscores permitted between
const validPattern = "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z]$"

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

func FirstValidName(strings ...string) (n Name, ok bool) {
	for _, s := range strings {
		n, ok = NewName(s)
		if ok {
			return
		}
	}

	return
}

func (n Name) String() string {
	return string(n)
}

func (n Name) Filename() string {
	return fmt.Sprintf("%s.%s", n, extension)
}
