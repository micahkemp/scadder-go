package renderer

import "testing"

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
		r := Renderer{OutputPath: test.path}

		filepath := r.filepath(test.filename)

		if filepath != test.filepath {
			t.Errorf("path %s filepath(%q) = %s, want %s", test.path, test.filename, filepath, test.filepath)
		}
	}
}
