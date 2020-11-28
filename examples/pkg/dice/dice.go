package dice

import "github.com/micahkemp/scad/pkg/scad"

type Dice struct {
	Name   string
	Size   float64
	Offset float64
}

var ExampleDice = Dice{
	Name: "example_dice",
	Size: 10,
}

func (d Dice) Render(path string) {
	die1 := Die{
		Name: "die1",
		Size: d.Size,
	}

	die2 := scad.Renderable(Die{
		Size: d.Size,
	}).TranslateWithName("die2", d.Size, 0, 0)

	name := scad.Name{
		Specified: d.Name,
		Default:   "dice",
	}

	dice := scad.Renderable(die1).AddWithName(name.String(), die2)

	dice.Render(".")
}
