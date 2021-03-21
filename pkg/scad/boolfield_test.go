package scad

import (
	"testing"
)

// BoolField's value should be set if created with NewDefinedBool, but not otherwise.
func TestBoolField_Defined(t *testing.T) {
	tests := []struct {
		input       BoolField
		wantDefined bool
	}{
		{
			BoolField{Name: "test", Value: ExplicitBool(true)},
			true,
		},
		{
			BoolField{Name: "test", Value: ExplicitBool(false)},
			true,
		},
		{
			BoolField{Name: "test"},
			false,
		},
	}

	for _, test := range tests {
		gotDefined := test.input.Value.isSet()
		if gotDefined != test.wantDefined {
			t.Errorf("%T{%+v}.Value.isSet() = %v", test.input, test.input, gotDefined)
		}
	}
}
