package scad

import (
	"fmt"
	"testing"
)

// fieldFormatterTest defines an input fieldFormatter and expected value.
type fieldFormatterTest struct {
	input fieldFormatter
	want  string
}

// run performs the test defined in a fieldFormatterTest.
func (f fieldFormatterTest) run(t *testing.T) {
	message := fmt.Sprintf("%T{%+v}.fieldFormat() = %q, want %q", f.input, f.input, f.input.fieldFormat(), f.want)

	if f.input.fieldFormat() == f.want {
		t.Log(message)
	} else {
		t.Errorf("!!!FAIL!!! %s", message)
	}
}
