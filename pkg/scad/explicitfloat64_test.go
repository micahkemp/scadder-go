package scad

import "testing"

func TestExplicitFloat64_SCADString(t *testing.T) {
	tests := scadStringerTests{
		{
			ExplicitFloat64(1),
			"1",
		},
		{
			ExplicitFloat64(1.001),
			"1.001",
		},
	}

	for _, test := range tests {
		if test.input.SCADString() != test.want {
			t.Errorf("%T{%+v}.SCADString() = %q, want %q", test.input, test.input, test.input.SCADString(), test.want)
		}
	}
}
