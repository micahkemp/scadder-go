package scad

import "testing"

func TestFields_String(t *testing.T) {
	tests := []struct {
		input fields
		want  string
	}{
		// vanilla, single field
		{
			fields{
				"fieldA": "valueA",
			},
			`fieldA=valueA`,
		},
		// quoted value
		{
			fields{
				"fieldA": `"valueA"`,
			},
			`fieldA="valueA"`,
		},
		// embedded quotes
		{
			fields{
				"fieldA": `value"A"`,
			},
			`fieldA=value"A"`,
		},
		// multiple fields, sorted by key
		{
			fields{
				"fieldC": "valueC",
				"fieldA": "valueA",
				"fieldB": "valueB",
			},
			`fieldA=valueA, fieldB=valueB, fieldC=valueC`,
		},
	}

	for _, test := range tests {
		if test.input.String() != test.want {
			t.Errorf("fields{%v}.String() = %q, want %q",
				test.input,
				test.input.String(),
				test.want)
		}
	}
}
