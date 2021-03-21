package scad

import (
	"fmt"
	"testing"
)

// fieldValueFormatterTest defines a test for a fieldValueFormatter.
type fieldValueFormatterTest struct {
	input FieldValueFormatter
	want  string
}

// run performs the test defined by a fieldValueFormatterTest.
func (s fieldValueFormatterTest) run(t *testing.T) {
	message := fmt.Sprintf("%T{%+v}.fieldValueFormat() = %q, want %q", s.input, s.input, s.input.fieldValueFormat(), s.want)
	if s.input.fieldValueFormat() == s.want {
		t.Log(message)
	} else {
		t.Errorf("!!!FAIL!!! %s", message)

	}
}
