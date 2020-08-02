package scad

import "strconv"

func ShortFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
