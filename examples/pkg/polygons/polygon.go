package polygons

import "github.com/micahkemp/scad/pkg/scad"

var ExamplePolygon = scad.Polygon{
	Name: "example_polygon",
	Points: scad.XYCoordinates{
		{10, 0},
		{10, 10},
		{-10, 0},
		{0, -10},
	},
}
