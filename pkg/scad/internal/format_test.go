package internal

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

func TestShortFloatList(t *testing.T) {
	tests := []struct {
		floats []float64
		want   string
	}{
		{[]float64{1}, "[1]"},
		{[]float64{1, 2.1}, "[1, 2.1]"},
		{[]float64{1, 2.1, 3.12}, "[1, 2.1, 3.12]"},
	}

	for _, test := range tests {
		got := ShortFloatList(test.floats...)

		if got != test.want {
			t.Errorf("ShortFloatList(%v) = \"%s\", want \"%s\"", test.floats, got, test.want)
		}
	}
}
