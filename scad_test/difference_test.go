package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestDifference(t *testing.T) {
	child1 := scad.NewCube("child1", 1, 2, 3, false)
	child2 := scad.NewCube("child2", 1, 2, 3, false)

	test := struct {
		name string
		want string
	}{
		"", `use <child1/child1.scad>
use <child2/child2.scad>

module difference_component() {
	difference() {
		child1();
		child2();
	}
}

// call module when loaded directly
difference_component();
`,
	}

	newDifference := scad.NewDifference(test.name, child1, child2)

	if newDifference.String() != test.want {
		t.Errorf("NewDifference(%q, child1, child2).String() = \n%s, want \n%s",
			test.name,
			newDifference.String(),
			test.want,
		)
	}

	// this should give us the same results as creating the Difference object directly
	// but we want to give a good failure message, so it's broken out separately
	child1Difference := child1.Subtract(test.name, child2)
	if child1Difference.String() != test.want {
		t.Errorf("child1.Subtract(%q, child2).String() = %s, want %s",
			test.name,
			child1Difference.String(),
			test.want,
		)
	}
}
