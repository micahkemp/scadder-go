package scad

import (
	"testing"
)

func TestNew(t *testing.T) {
	testName, _ := NewName("testName")
	callName := "callModule"

	tests := []struct {
		template string
		fields   Fields
		want     string
	}{
		{"no variables", Fields{}, "no variables"},
		{"{{ .Name.String }}", Fields{}, testName.String()},
		{"{{ .CallName }}", Fields{}, callName},
		{"{{ .Fields.String }}", NewFields(map[string]string{"fieldA": "valueA"}), "fieldA=valueA"},
	}

	for _, test := range tests {
		template := NewTemplate(testName, test.template, callName, test.fields)
		got := template.String()
		if got != test.want {
			t.Errorf("New(%v, %v).String() = %s, want %s",
				test.template,
				test.fields,
				got,
				test.want)
		}
	}
}
