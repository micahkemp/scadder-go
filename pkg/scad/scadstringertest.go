package scad

import "testing"

// scadStringerTest defines a test for a scadStringer.
type scadStringerTest struct {
	input scadStringer
	want  string
}

// run performs the test defined by a scadStringerTest.
func (s scadStringerTest) run(t *testing.T) {
	if s.input.SCADString() != s.want {
		t.Errorf("%T{%+v}.SCADString() = %q, want %q", s.input, s.input, s.input.SCADString(), s.want)
	}
}
