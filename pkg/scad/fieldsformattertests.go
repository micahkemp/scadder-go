package scad

import "testing"

// fieldsFormatterTests is a list of fieldsFormatterTest objects.
type fieldsFormatterTests []fieldsFormatterTest

// run calls run for each of a fieldsFormatterTests members.
func (f fieldsFormatterTests) run(t *testing.T) {
	for _, test := range f {
		test.run(t)
	}
}
