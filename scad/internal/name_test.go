package internal

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

func TestFilenameFilePath(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		wantFilename string
		wantFilePath string
	}{
		{"aa", ".", "aa.scad", "aa.scad"},
		{"bb", "relative_output_path", "bb.scad", "relative_output_path/bb.scad"},
		{"cc", "/absolute_output_path", "cc.scad", "/absolute_output_path/cc.scad"},
	}

	for _, test := range tests {
		n, _ := NewName(test.name)
		gotFilename := n.Filename()
		gotFilePath := n.FilePath(test.path)

		if gotFilename != test.wantFilename {
			t.Errorf("NewName(%q).Filename() = %q, want %q",
				test.name,
				gotFilename,
				test.wantFilename)
		}

		if gotFilePath != test.wantFilePath {
			t.Errorf("NewName(%q).FilePath(%q) = %q, want %q",
				test.name,
				test.path,
				gotFilePath,
				test.wantFilePath)
		}
	}
}
