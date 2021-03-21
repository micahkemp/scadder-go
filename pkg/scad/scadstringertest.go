package scad

import (
	"fmt"
	"testing"
)

// scadStringerTest defines a test for a scadStringer.
type scadStringerTest struct {
	input isSetSCADStringer
	want  string
}

// run performs the test defined by a scadStringerTest.
func (s scadStringerTest) run(t *testing.T) {
	message := fmt.Sprintf("%T{%+v}.scadString() = %q, want %q", s.input, s.input, s.input.scadString(), s.want)
	if s.input.scadString() == s.want {
		t.Log(message)
	} else {
		t.Errorf("!!!FAIL!!! %s", message)

	}
}
