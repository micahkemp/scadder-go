package template

import (
	"github.com/micahkemp/scad/fields"
	"github.com/micahkemp/scad/name"
	"testing"
)

func TestNew(t *testing.T) {
	testName, _ := name.New("testName")

	tests := []struct {
		template string
		fields   fields.Fields
		want     string
	}{
		{"no variables", fields.Fields{}, "no variables"},
		{"{{ .Name.String }}", fields.Fields{}, testName.String()},
		{"{{ .Fields.String }}", fields.New(map[string]string{"fieldA": "valueA"}), "fieldA=valueA"},
	}

	for _, test := range tests {
		template := New(testName, test.template, test.fields)
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
