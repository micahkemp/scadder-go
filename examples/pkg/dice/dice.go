package dice

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad"
)

type Dice struct {
	Name   string
	Count  int
	Die    Die
	Offset float64
}

var ExampleDice = Dice{
	Name:  "example_dice",
	Count: 5,
	Die: Die{
		Size: 50,
		Dimple: Dimple{
			Diameter: 10,
			Depth:    6,
		},
	},
	Offset: 5,
}

func (d Dice) SCADWriter() scad.SCADWriter {
	translatedDice := make([]scad.Module, d.Count)

	for dieNumber, _ := range translatedDice {
		translatedDice[dieNumber] = scad.TranslationWithName(
			fmt.Sprintf("die_%d", dieNumber),
			float64(dieNumber)*(d.Die.Size+d.Offset),
			0,
			0,
			d.Die,
		)
	}

	name := scad.Name{
		Specified: d.Name,
		Default:   "dice",
	}

	dice := scad.UnionWithName(name.String(), translatedDice...)

	return dice.SCADWriter()
}
