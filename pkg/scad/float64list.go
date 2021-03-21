package scad

import (
	"fmt"
	"strings"
)

// Float64List is a list of float64 values.
type Float64List []float64

// isSet returns true if the list has more than zero items.
func (f Float64List) isSet() bool {
	return len(f) > 0
}

// SCADString returns Float64List as a string suitable to use in OpenSCAD.
func (f Float64List) SCADString() string {
	valueStrings := make([]string, len(f))

	for i, value := range f {
		iExplicitFloat64 := ExplicitFloat64(value)
		valueStrings[i] = iExplicitFloat64.SCADString()
	}

	allValuesString := strings.Join(valueStrings, ", ")
	return fmt.Sprintf("[%s]", allValuesString)
}
