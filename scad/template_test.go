package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"testing"
)

func TestNew(t *testing.T) {
	testName, _ := internal.NewName("testName")
	callName := "callModule"

	tests := []struct {
		template string
		fields   internal.Fields
		want     string
	}{
		{"no variables", internal.Fields{}, "no variables"},
		{"{{ .Name.String }}", internal.Fields{}, testName.String()},
		{"{{ .CallName }}", internal.Fields{}, callName},
		{"{{ .Fields.String }}", internal.NewFields(map[string]string{"fieldA": "valueA"}), "fieldA=valueA"},
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
