package internal

import (
	"fmt"
	"sort"
	"strings"
)

type Fields struct {
	values map[string]string
}

func NewFields(v map[string]string) Fields {
	return Fields{v}
}

func (f Fields) String() string {
	var keys []string
	for k := range f.values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var keyValues []string
	for _, k := range keys {
		keyValues = append(keyValues, fmt.Sprintf("%s=%s", k, f.values[k]))
	}

	return strings.Join(keyValues, ", ")
}
