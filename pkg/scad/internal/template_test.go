package internal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	testName := "testName"
	callName := "callModule"

	tests := []struct {
		template string
		fields   Fields
		want     string
	}{
		{"no variables", Fields{}, "no variables"},
		{"{{ .Name.String }}", Fields{}, testName},
		{"{{ .CallName }}", Fields{}, callName},
		{"{{ .Fields.String }}", NewFields(map[string]string{"fieldA": "valueA"}), "fieldA=valueA"},
	}

	for _, test := range tests {
		template := newTemplate(testName, test.template, callName, test.fields)
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

func TestRendered(t *testing.T) {
	child := newTemplate("child_template", "template contents", "call_name", Fields{})
	parent := newTemplate("parent_template", "template contents", "call_name", Fields{})
	parent.Children = []DirPathRenderer{child}
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
