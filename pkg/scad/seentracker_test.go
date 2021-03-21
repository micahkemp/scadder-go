package scad

import "testing"

func TestWriteTracker_trackPath(t *testing.T) {
	tests := []struct {
		input    string
		wantSeen bool
	}{
		{"firstPath", true},
		{"secondPath", true},
		{"firstPath", false},
		{"secondPath", false},
		{"thirdPath", true},
	}

	s := seenTracker{}

	for _, test := range tests {
		gotSeen := s.trackPath(test.input)
		if gotSeen != test.wantSeen {
			t.Errorf("%T{%+v}.trackPath(%q) = %v, want %v", s, s, test.input, gotSeen, test.wantSeen)
		}
	}
}
