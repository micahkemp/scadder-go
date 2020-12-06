package cylinders

import "github.com/micahkemp/scad/pkg/scad"

var ExampleCylinderRadius = scad.Cylinder{
	Name:   "cylinder_radius",
	Height: 20,
	Radius: 5,
}

var ExampleCylinderR1R2 = scad.Cylinder{
	Name:    "cylinder_r1_r2",
	Height:  20,
	Radius1: 10,
	Radius2: 0,
}

var ExampleCylinderDiameter = scad.Cylinder{
	Name:     "cylinder_diameter",
	Height:   20,
	Diameter: 5,
}

var ExampleCylinderD1D2 = scad.Cylinder{
	Name:      "cylinder_d1_d2",
	Height:    20,
	Diameter1: 10,
	Diameter2: 0,
}
