package name

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a", false},
		{"0a", false},
		{"a0", false},
		{"a-a", false},
		{"_a", false},
		{"a_", false},
		{"a a", false},
		{"a0a", true},
		{"aa", true},
		{"a_a", true},
	}

	for _, test := range tests {
		if got := valid(test.input); got != test.want {
			t.Errorf("Valid(%q) = %v", test.input, got)
		}
	}
}
