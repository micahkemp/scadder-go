package scad

import (
	"fmt"
	"strings"
)

type XYCoordinate struct {
	X, Y float64
}

type XYCoordinates []XYCoordinate

type XYZCoordinate struct {
	X, Y, Z float64
}

func (p XYCoordinate) String() string {
	return shortFloatList(p.X, p.Y)
}

func (p XYCoordinates) String() string {
	pointStrings := make([]string, len(p))
	for pointNumber, point := range p {
		pointStrings[pointNumber] = point.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(pointStrings, ", "))
}

func (p XYZCoordinate) String() string {
	return shortFloatList(p.X, p.Y, p.Z)
}
