package scad

import "testing"

func TestName_Value(t *testing.T) {
	tests := []struct {
		specifiedValue string
		defaultValue   string
		wantValue      string
		wantOK         bool
	}{
		{"", "", "", false},
		{"", "default", "default", true},
		{"specified", "default", "specified", true},
		{"specified", "", "specified", true},
	}

	for _, test := range tests {
		name := Name{
			specifiedValue: test.specifiedValue,
			defaultValue:   test.defaultValue,
		}

		value, ok := name.value()

		if value != test.wantValue || ok != test.wantOK {
			t.Errorf("Name{%q, %q}.value() = (%q, %v), want (%q, %v)",
				test.specifiedValue,
				test.defaultValue,
				value,
				ok,
				test.wantValue,
				test.wantOK)
		}
	}
}
