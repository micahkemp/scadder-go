package scad

import (
	"fmt"
	"testing"
)

// fieldsformatterTest defines a test case for a fieldsFormater object.
type fieldsFormatterTest struct {
	input fieldsFormatter
	want  string
}

// run performs the test defined by a fieldsFormatterTest.
func (f fieldsFormatterTest) run(t *testing.T) {
	message := fmt.Sprintf("%T{%+v}.fieldsFormat() = %q, want %q", f.input, f.input, f.input.fieldsFormat(), f.want)

	if f.input.fieldsFormat() == f.want {
		t.Log(message)
	} else {
		t.Errorf("!!!FAIL!!! %s", message)
	}
}
