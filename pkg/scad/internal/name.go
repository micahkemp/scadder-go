package internal

import "log"

type Name struct {
	Specified string
	Default   string
}

func (n Name) Value() (value string, ok bool) {
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
	value, ok := n.Value()
	if !ok {
		log.Fatal("Attempted to stringify empty Name")
	}

	return value
}
