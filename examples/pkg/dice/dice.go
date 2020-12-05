package dice

import "github.com/micahkemp/scad/pkg/scad"

type Dice struct {
	Name   string
	Size   float64
	Offset float64
}

var ExampleDice = Dice{
	Name:   "example_dice",
	Size:   10,
	Offset: 5,
}

func (d Dice) SCADWriter() scad.SCADWriter {
	die1 := Die{
		Name: "die1",
		Size: d.Size,
	}

	die2 := scad.TranslationWithName(
		"die2",
		d.Size+d.Offset,
		0,
		0,
		Die{Size: d.Size},
	)

	name := scad.Name{
		Specified: d.Name,
		Default:   "dice",
	}

	dice := scad.UnionWithName(name.String(), die1, die2)

	return dice.SCADWriter()
}
