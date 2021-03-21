package scad

import "testing"

func TestField_SCADString(t *testing.T) {
	tests := fieldFormatterTests{
		{
			field{
				"fieldName",
				Float64List{1, 0.1, 0.01},
			},
			"fieldName=[1, 0.1, 0.01]",
		},
		{
			field{
				"fieldName",
				ExplicitBool(false),
			},
			"fieldName=false",
		},
	}

	tests.run(t)
}
