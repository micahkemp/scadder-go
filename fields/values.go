package fields

import (
	"fmt"
	"sort"
	"strings"
)

func (f Fields) String() string {
	var keys []string
	for k := range f.values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var keyValues []string
	for _, k := range keys {
		keyValues = append(keyValues, fmt.Sprintf("%s=%q", k, f.values[k]))
	}

	return strings.Join(keyValues, ", ")
}
