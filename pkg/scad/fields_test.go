package scad

import "testing"

func TestFields_fieldsFormat(t *testing.T) {
	tests := fieldsFormatterTests{
		{
			fields{
				field{name: "includedField", FieldValueFormatter: ExplicitBool(false)},
			},
			"includedField=false",
		},
		{
			fields{
				field{name: "skippedField", FieldValueFormatter: BoolWithExplicitness{}},
				field{name: "includedField1", FieldValueFormatter: ExplicitBool(true)},
				field{name: "includedField2", FieldValueFormatter: ExplicitFloat64(1)},
			},
			"includedField1=true, includedField2=1",
		},
	}

	tests.run(t)
}
