package scad

import "testing"

func TestFloat64List_scadString(t *testing.T) {
	tests := fieldValueFormatterTests{
		{
			Float64List{1.0, 0.1, 0.01},
			"[1, 0.1, 0.01]",
		},
	}

	tests.run(t)
}
