package scadmodule

import (
	"testing"
)

type hasName struct {name string}
func (h hasName) Name() string { return h.name }

type hasEmbeddedName struct {
	*hasName
}

func TestHasName(t *testing.T) {
	tests := []struct {
		set		string
		want 	string
	}{
		{"", "hasName"},
		{"overridden", "overridden"},
	}

	for _, test := range(tests) {
		h := hasName{
			name: test.set,
		}
		if selectName(h) != test.want {
			t.Errorf("selectName(hasName{name: %q}) = %s, want %s",
				test.set,
				selectName(h),
				test.want)
		}
	}
}

func TestHasEmbeddedName(t *testing.T) {
	tests := []struct {
		set		string
		want	string
	}{
		{"", "hasEmbeddedName"},
		{"overridden", "overridden"},
	}

	for _, test := range tests {
		h := hasEmbeddedName{
			&hasName{name: test.set},
		}
		if selectName(h) != test.want {
			t.Errorf("selectName(hasEmbeddedName{&hasName{name: %q}}) = %s, want %s",
				test.set,
				selectName(h),
				test.want)
		}
	}
}