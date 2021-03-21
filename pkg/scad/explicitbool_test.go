package scad

import "testing"

func TestExplicitBool_isSet(t *testing.T) {
	tests := []struct {
		input BoolWithExplicitness
		want  bool
	}{
		{
			BoolWithExplicitness{},
			false,
		},
		{
			BoolWithExplicitness{false, true},
			true,
		},
	}

	for _, test := range tests {
		if test.input.isSet() != test.want {
			t.Errorf("%T{%+v}.isSet() = %v, want %v", test.input, test.input, test.input.isSet(), test.want)
		}
	}
}
