package scadmodule

import "testing"

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
