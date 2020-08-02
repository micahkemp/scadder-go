package scad

import (
	"fmt"
	"strconv"
	"strings"
)

func ShortFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func ShortFloatList(floats ...float64) string {
	var shortFloats []string

	for _, f := range floats {
		shortFloats = append(shortFloats, ShortFloat(f))
	}

	return fmt.Sprintf("[%s]", strings.Join(shortFloats, ", "))
}
