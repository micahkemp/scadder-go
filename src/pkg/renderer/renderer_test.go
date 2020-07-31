package renderer

import "testing"

func TestRenderer(t *testing.T) {
	tests := []struct {
		t  string
		n  string
		ok bool
	}{
		{"", "", false},
		{"template", "", false},
		{"", "name", false},
		{"template", "test", true},
	}

	for _, test := range tests {
		r := Renderer{test.t, test.n}
		ok := r.Render("", *new(Fields))
		if ok != test.ok {
			t.Errorf("%v.render() = %v", r, ok)
		}
	}
}
