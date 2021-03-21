package scad

import "testing"

// scadStringerTests is a list of scadStringerTest objects.
type scadStringerTests []scadStringerTest

// run calls run for each of the scadStringerTests members.
func (s scadStringerTests) run(t *testing.T) {
	for _, test := range s {
		test.run(t)
	}
}
