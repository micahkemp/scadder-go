package scad

import "testing"

func TestName_Value(t *testing.T) {
	tests := []struct {
		Specified string
		Default   string
		WantValue string
		WantOK    bool
	}{
		{"", "", "", false},
		{"", "default", "default", true},
		{"specified", "default", "specified", true},
		{"specified", "", "specified", true},
	}

	for _, test := range tests {
		name := Name{
			Specified: test.Specified,
			Default:   test.Default,
		}

		value, ok := name.value()

		if value != test.WantValue || ok != test.WantOK {
			t.Errorf("Name{%q, %q}.value() = (%q, %v), want (%q, %v)",
				test.Specified,
				test.Default,
				value,
				ok,
				test.WantValue,
				test.WantOK)
		}
	}
}
