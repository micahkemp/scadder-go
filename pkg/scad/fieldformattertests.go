package scad

import "testing"

// fieldFormatterTests is a list of fieldFormatterTest objects.
type fieldFormatterTests []fieldFormatterTest

// run calls run for each of a fieldFormatterTests members.
func (f fieldFormatterTests) run(t *testing.T) {
	for _, test := range f {
		test.run(t)
	}
}
