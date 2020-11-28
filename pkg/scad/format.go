package scad

import (
	"fmt"
	"strconv"
	"strings"
)

func shortFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func shortFloatList(floats ...float64) string {
	var shortFloats []string

	for _, f := range floats {
		shortFloats = append(shortFloats, shortFloat(f))
	}

	return fmt.Sprintf("[%s]", strings.Join(shortFloats, ", "))
}
