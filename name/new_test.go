package name

import "testing"

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
		if got, ok := FirstValid(test.choices...); got != Name(test.wantName) || ok != test.wantOk {
			t.Errorf("FirstValid(%q) = (%q, %v), want (%q, %v)",
				test.choices,
				got,
				ok,
				test.wantName,
				test.wantOk)
		}
	}
}
