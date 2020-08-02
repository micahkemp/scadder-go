package fields

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		input map[string]string
		want  string
	}{
		// vanilla, single field
		{map[string]string{
			"fieldA": "valueA",
		},
			`fieldA=valueA`,
		},
		// quoted value
		{map[string]string{
			"fieldA": `"valueA"`,
		},
			`fieldA="valueA"`,
		},
		// embedded quotes
		{map[string]string{
			"fieldA": `value"A"`,
		},
			`fieldA=value"A"`,
		},
		// multiple fields, sorted by key
		{map[string]string{
			"fieldC": "valueC",
			"fieldA": "valueA",
			"fieldB": "valueB",
		},
			`fieldA=valueA, fieldB=valueB, fieldC=valueC`,
		},
	}

	for _, test := range tests {
		f := NewFields(test.input)

		if f.String() != test.want {
			t.Errorf("New(%q).String() = %s, want %s",
				test.input,
				f.String(),
				test.want)
		}
	}
}
