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

func TestChildPath(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		childPath string
	}{
		{"testA", ".", "testA"},
		{"testB", "output_path", "output_path/testB"},
	}

	// children don't currently have to be children of the parent for this test to succeed
	// that's ok
	parent := NewTemplate("parent", "", "", internal.Fields{})

	for _, test := range tests {
		childName, _ := internal.NewName(test.name)
		child := NewTemplate(childName, "", "", internal.Fields{})
		childPath := parent.childPath(child, test.path)

		if childPath != test.childPath {
			t.Errorf("parent.childPath(%v, %q) = %q, want %q",
				child,
				test.path,
				childPath,
				test.childPath)
		}
	}

}

func TestRendered(t *testing.T) {
	template := NewTemplate("test_template", "template contents", "call_name", internal.Fields{})
	renderPath := "."

	if template.rendered(renderPath) {
		t.Errorf("before rendering, template.rendered(%q) = %v",
			renderPath,
			template.rendered(renderPath))
	}

	template.Render(renderPath)

	if !template.rendered(renderPath) {
		t.Errorf("after rendering, template.rendered(%q) = %v",
			renderPath,
			template.rendered(renderPath))
	}
}
