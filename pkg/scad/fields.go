package scad

import (
	"fmt"
	"sort"
	"strings"
)

type fields map[string]string

func (f fields) String() string {
	var keys []string
	for k := range f {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var keyValues []string
	for _, k := range keys {
		keyValues = append(keyValues, fmt.Sprintf("%s=%s", k, f[k]))
	}

	return strings.Join(keyValues, ", ")
}
