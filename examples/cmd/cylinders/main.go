package main

import (
	"github.com/micahkemp/scad/examples/pkg/cylinders"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	scad.CLIHandler(cylinders.Models)
}
