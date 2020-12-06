package cylinders

import "github.com/micahkemp/scad/pkg/scad"

var Models = scad.Models{
	"cylinder_radius":   &ExampleCylinderRadius,
	"cylinder_r1_r2":    &ExampleCylinderR1R2,
	"cylinder_diameter": &ExampleCylinderDiameter,
	"cylinder_d1_d2":    &ExampleCylinderD1D2,
}
