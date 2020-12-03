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

var DimplePositions1 = DimplePositions{
	{0, 0},
}

var DimplePositions2 = DimplePositions{
	{-0.33, -0.33},
	{0.33, 0.33},
}

var DimplePositions3 = DimplePositions{
	{-0.33, -0.33},
	{0, 0},
	{0.33, 0.33},
}

var DimplePositions4 = DimplePositions{
	{-0.33, -0.33},
	{-0.33, 0.33},
	{0.33, 0.33},
	{0.33, -0.33},
}

var DimplePositions5 = DimplePositions{
	{-0.33, -0.33},
	{-0.33, 0.33},
	{0, 0},
	{0.33, 0.33},
	{0.33, -0.33},
}

var DimplePositions6 = DimplePositions{
	{-0.33, -0.33},
	{-0.33, 0},
	{-0.33, 0.33},
	{0.33, 0.33},
	{0.33, 0},
	{0.33, -0.33},
}

func exampleDimplesWithPositions(p DimplePositions) Dimples {
	return Dimples{
		Dimple: Dimple{
			Diameter: 10,
			Depth:    8,
		},
		Length:          50,
		DimplePositions: p,
	}
}

var ExampleDimples1 = exampleDimplesWithPositions(DimplePositions1)
var ExampleDimples2 = exampleDimplesWithPositions(DimplePositions2)
var ExampleDimples3 = exampleDimplesWithPositions(DimplePositions3)
var ExampleDimples4 = exampleDimplesWithPositions(DimplePositions4)
var ExampleDimples5 = exampleDimplesWithPositions(DimplePositions5)
var ExampleDimples6 = exampleDimplesWithPositions(DimplePositions6)

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
