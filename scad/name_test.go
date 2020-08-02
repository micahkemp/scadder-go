package scad

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

func TestFirstValid(t *testing.T) {
	tests := []struct {
		choices  []string
		wantName string
		wantOk   bool
	}{
		{[]string{"", "a"}, "", false},
		{[]string{"", "a", "aa", "bb", "cc"}, "aa", true},
	}

	for _, test := range tests {
		if got, ok := FirstValidName(test.choices...); got != Name(test.wantName) || ok != test.wantOk {
			t.Errorf("FirstValidName(%q) = (%q, %v), want (%q, %v)",
				test.choices,
				got,
				ok,
				test.wantName,
				test.wantOk)
		}
	}
}

func TestFilename(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"aa", "aa.scad"},
		{"bb", "bb.scad"},
	}

	for _, test := range tests {
		n, _ := NewName(test.input)
		got := n.Filename()

		if got != test.want {
			t.Errorf("NewName(%q).Filename() = %q, want %q",
				test.input,
				got,
				test.want)
		}
	}
}
