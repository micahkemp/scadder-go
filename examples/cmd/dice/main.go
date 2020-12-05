package main

import (
	"github.com/micahkemp/scad/examples/pkg/dice"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	scad.CLIHandler(dice.Models)
}
