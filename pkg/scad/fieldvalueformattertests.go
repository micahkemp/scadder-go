package scad

import "testing"

// fieldValueFormatterTests is a list of fieldValueFormatterTest objects.
type fieldValueFormatterTests []fieldValueFormatterTest

// run calls run for each of the scadStringerTests members.
func (s fieldValueFormatterTests) run(t *testing.T) {
	for _, test := range s {
		test.run(t)
	}
}
