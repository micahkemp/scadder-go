package scadmodule

import (
	"testing"
)

func TestFilename(t *testing.T) {
	var tests = []struct {
		name     string
		filename string
		ok       bool
	}{
		{"", ".scad", false},
		{"a b", "a b.scad", false},
		{"1a", "1a.scad", false},
		{"a-1", "a-1.scad", false},
		{"a", "a.scad", true},
		{"a1", "a1.scad", true},
		{"a_1", "a_1.scad", true},
	}

	for _, test := range tests {
		r := SCADModule{
			Name: test.name,
		}

		filename, ok := r.filename()
		if filename != test.filename {
			t.Errorf("Render{Name: %s}.filename() filename = %s, want %s",
				test.name,
				filename,
				test.filename)
		}

		if ok != test.ok {
			t.Errorf("Render{Name: %s}.filname() ok = %v", test.name, ok)
		}
	}
}

func TestFilepath(t *testing.T) {
	tests := []struct {
		path     string
		filename string
		filepath string
	}{
		{"", "ok.scad", "ok.scad"},
		{".", "ok.scad", "ok.scad"},
		{"output", "ok.scad", "output/ok.scad"},
	}

	for _, test := range tests {
		s := SCADModule{}

		filepath := s.filepath(test.path, test.filename)

		if filepath != test.filepath {
			t.Errorf("path %s filepath(%q) = %s, want %s", test.path, test.filename, filepath, test.filepath)
		}
	}
}

func TestSCADModuleIsRenderable(t *testing.T) {
	var _ Renderer = SCADModule{}
}
