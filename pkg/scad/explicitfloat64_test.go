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

	tests.run(t)
}
