package main

import (
	"github.com/micahkemp/scad/examples/pkg/polygons"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	scad.CLIHandler(polygons.Models)
}
