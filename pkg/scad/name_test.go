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
		{"a-a", false},
		{"_a", false},
		{"a_", false},
		{"a a", false},
		{"a0a", true},
		{"aa", true},
		{"a_a", true},
		{"a0", true},
	}

	for _, test := range tests {
		n := name{SpecifiedName: test.input}
		if got := n.Valid(); got != test.want {
			t.Errorf("Valid(%q) = %v", test.input, got)
		}
	}

	n := name{}
	if got := n.Valid(); got {
		t.Errorf("Name{}.Valid() = %v", got)
	}
}

func TestDefaultOrSpecified(t *testing.T) {
	tests := []struct {
		name      name
		wantValue string
	}{
		{name{"", "a"}, "a"},
		{name{"a", ""}, "a"},
		{name{"", ""}, ""},
	}

	for _, test := range tests {
		if test.name.Value() != test.wantValue {
			t.Errorf("%q.Value() = %s, want %s",
				test.name,
				test.name.Value(),
				test.wantValue)

		}
	}
}

func TestFilenameFilePath(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		wantDirPath  string
		wantFilename string
		wantFilePath string
	}{
		{"aa", ".", "aa", "aa.scad", "aa.scad"},
		{"bb", "relative_output_path", "relative_output_path/bb", "bb.scad", "relative_output_path/bb.scad"},
		{"cc", "/absolute_output_path", "/absolute_output_path/cc", "cc.scad", "/absolute_output_path/cc.scad"},
	}

	for _, test := range tests {
		n := name{defaultName: test.name}
		gotDirPath := n.DirPath(test.path)
		gotFilename := n.Filename()
		gotFilePath := n.FilePath(test.path)

		if gotDirPath != test.wantDirPath {
			t.Errorf("NewName(%q).DirPath(%q) = #{%q}, want %q",
				test.name,
				test.path,
				gotDirPath,
				test.wantDirPath)
		}

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
