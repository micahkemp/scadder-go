package scad

import (
	"fmt"
	"github.com/micahkemp/scad/scad/internal"
	"strings"
)

type Point2D struct{ X, Y float64 }

func (p Point2D) String() string {
	return internal.ShortFloatList(p.X, p.Y)
}

type Points2D []Point2D

func (points Points2D) String() string {
	stringPoints := make([]string, len(points))
	for i, point := range points {
		stringPoint := point.String()
		stringPoints[i] = stringPoint
	}

	return fmt.Sprintf("[%s]", strings.Join(stringPoints, ", "))
}
