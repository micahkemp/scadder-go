package template

import (
	"github.com/micahkemp/scad/fields"
	"github.com/micahkemp/scad/name"
	"testing"
)

func TestNew(t *testing.T) {
	testName, _ := name.NewName("testName")
	callName := "callModule"

	tests := []struct {
		template string
		fields   fields.Fields
		want     string
	}{
		{"no variables", fields.Fields{}, "no variables"},
		{"{{ .Name.String }}", fields.Fields{}, testName.String()},
		{"{{ .CallName }}", fields.Fields{}, callName},
		{"{{ .Fields.String }}", fields.NewFields(map[string]string{"fieldA": "valueA"}), "fieldA=valueA"},
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
