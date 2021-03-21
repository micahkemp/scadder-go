package scad

import "testing"

func TestFields_scadString(t *testing.T) {
	tests := scadStringerTests{
		{
			fields{
				field{name: "includedField", isSetSCADStringer: ExplicitBool(false)},
			},
			"includedField=false",
		},
		{
			fields{
				field{name: "skippedField", isSetSCADStringer: BoolWithExplicitness{}},
				field{name: "includedField1", isSetSCADStringer: ExplicitBool(true)},
				field{name: "includedField2", isSetSCADStringer: ExplicitFloat64(1)},
			},
			"includedField1=true, includedField2=1",
		},
	}

	tests.run(t)
}
