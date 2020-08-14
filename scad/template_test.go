package scad

import (
	"github.com/micahkemp/scad/scad/internal"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	testName := "testName"
	callName := "callModule"

	tests := []struct {
		template string
		fields   internal.Fields
		want     string
	}{
		{"no variables", internal.Fields{}, "no variables"},
		{"{{ .Name.String }}", internal.Fields{}, testName},
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
		childName := test.name
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
	child := NewTemplate("child_template", "template contents", "call_name", internal.Fields{})
	parent := NewTemplate("parent_template", "template contents", "call_name", internal.Fields{})
	parent.Children = []scadTemplate{child}
	renderPath, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("ioutil.TempDir(%q, %q) err: %s",
			"",
			"",
			err)
	}
	defer os.RemoveAll(renderPath)

	if parent.rendered(renderPath) {
		t.Errorf("before rendering, parent.rendered(%q) = %v",
			renderPath,
			parent.rendered(renderPath))
	}

	if parent.childRendered(child, renderPath) {
		t.Errorf("before rendering, parent.childRendered(%q) = %v",
			renderPath,
			parent.childRendered(child, renderPath))
	}

	// render twice, both should come back as ok
	for i := 1; i <= 2; i++ {
		ok := parent.Render(renderPath)
		if !ok {
			t.Errorf("[i=%d] parent.Render(%q) = %v",
				i,
				renderPath,
				ok)
		}
	}

	if !parent.rendered(renderPath) {
		t.Errorf("after rendering, parent.rendered(%q) = %v",
			renderPath,
			parent.rendered(renderPath))
	}

	if !parent.childRendered(child, renderPath) {
		t.Errorf("after rendering, parent.childRendered(%q) = %v",
			renderPath,
			parent.childRendered(child, renderPath))
	}
}
