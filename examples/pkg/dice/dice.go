package dice

import "github.com/micahkemp/scad/pkg/scad"

type Dice struct {
	Size   float64
	Offset float64
}

var ExampleDice = Dice{
	Size: 10,
}

func (d Dice) Render(path string) {
	die1 := Die{
		Name: "die1",
		Size: d.Size,
	}

	die2 := Die{
		Name: "die2",
		Size: d.Size,
	}

	die2Translated := scad.Translate{
		Name: "die2_translated",
		X:    d.Size + d.Offset,
		Children: []scad.Renderer{
			die2,
		},
	}

	dice := scad.Union{
		Children: []scad.Renderer{
			die1,
			die2Translated,
		},
	}

	dice.Render(".")
}
