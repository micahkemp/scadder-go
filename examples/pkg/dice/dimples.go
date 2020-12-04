package dice

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad"
)

type DimplePositions []scad.XYPoint

type Dimples struct {
	Name string
	Dimple
	DimplePositions
	Length float64
}

var DimplePositionsForCount = []DimplePositions{
	// 0
	0: {},
	// 1
	1: {
		{0, 0},
	},
	// 2
	2: {
		{-0.33, -0.33},
		{0.33, 0.33},
	},
	// 3
	3: {
		{-0.33, -0.33},
		{0, 0},
		{0.33, 0.33},
	},
	// 4
	4: {
		{-0.33, -0.33},
		{-0.33, 0.33},
		{0.33, 0.33},
		{0.33, -0.33},
	},
	// 5
	5: {
		{-0.33, -0.33},
		{-0.33, 0.33},
		{0, 0},
		{0.33, 0.33},
		{0.33, -0.33},
	},
	// 6
	6: {
		{-0.33, -0.33},
		{-0.33, 0},
		{-0.33, 0.33},
		{0.33, 0.33},
		{0.33, 0},
		{0.33, -0.33},
	},
}

func exampleDimplesWithCount(count int) Dimples {
	return Dimples{
		Dimple: Dimple{
			Diameter: 10,
			Depth:    8,
		},
		Length:          50,
		DimplePositions: DimplePositionsForCount[count],
	}
}

var ExampleDimples1 = exampleDimplesWithCount(1)
var ExampleDimples2 = exampleDimplesWithCount(2)
var ExampleDimples3 = exampleDimplesWithCount(3)
var ExampleDimples4 = exampleDimplesWithCount(4)
var ExampleDimples5 = exampleDimplesWithCount(5)
var ExampleDimples6 = exampleDimplesWithCount(6)

func (d Dimples) SCADWriter() scad.SCADWriter {
	allDimples := make([]scad.Module, len(d.DimplePositions))

	for dimpleNumber, dimplePosition := range d.DimplePositions {
		allDimples[dimpleNumber] = scad.Transformable(d.Dimple).TranslateWithName(
			fmt.Sprintf("dimple_%d", dimpleNumber),
			d.Length*dimplePosition.X,
			d.Length*dimplePosition.Y,
			0,
		)
	}

	return scad.Union{
		Name: scad.Name{
			Specified: d.Name,
			Default:   "dimples",
		}.String(),
		Children: allDimples,
	}.SCADWriter()
}
