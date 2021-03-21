package scad

import "testing"

func TestSeenTracker_trackSeen(t *testing.T) {
	tests := []struct {
		input    string
		wantSeen bool
	}{
		{"firstPath", false},
		{"secondPath", false},
		{"firstPath", true},
		{"secondPath", true},
		{"thirdPath", false},
	}

	s := seenTracker{}

	for _, test := range tests {
		gotSeen := s.trackSeen(test.input)
		if gotSeen != test.wantSeen {
			t.Errorf("%T{%+v}.trackPath(%q) = %v, want %v", s, s, test.input, gotSeen, test.wantSeen)
		}
	}
}
