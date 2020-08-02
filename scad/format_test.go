package scad

import "testing"

func TestShortFloat(t *testing.T) {
	tests := []struct {
		input float64
		want  string
	}{
		{1, "1"},
		{1.0, "1"},
		{1.1, "1.1"},
		{1.23456, "1.23456"},
	}

	for _, test := range tests {
		got := ShortFloat(test.input)

		if got != test.want {
			t.Errorf("ShortFloat(%f) = %s, want %s", test.input, got, test.want)
		}
	}
}
