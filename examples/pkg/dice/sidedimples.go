package dice

import (
	"fmt"
	"github.com/micahkemp/scad/pkg/scad"
)

type SideTranslation struct {
	RotateVector    scad.XYZCoordinate
	RotateAngle     int
	TranslateFactor scad.XYZCoordinate
}

// no translation
var SideTranslationBottom = SideTranslation{
	// down half height
	TranslateFactor: scad.XYZCoordinate{Z: -0.5},
}

var SideTranslationTop = SideTranslation{
	// around X axis
	RotateVector: scad.XYZCoordinate{X: 1},
	RotateAngle:  180,
	// up half height
	TranslateFactor: scad.XYZCoordinate{Z: 0.5},
}

var SideTranslationLeft = SideTranslation{
	// around Y axis
	RotateVector: scad.XYZCoordinate{Y: 1},
	RotateAngle:  90,
	// left half height
	TranslateFactor: scad.XYZCoordinate{X: -0.5},
}

var SideTranslationRight = SideTranslation{
	// around Y axis, the other direction
	RotateVector: scad.XYZCoordinate{Y: -1},
	RotateAngle:  90,
	// right half height
	TranslateFactor: scad.XYZCoordinate{X: 0.5},
}

var SideTranslationFront = SideTranslation{
	// around X axis
	RotateVector: scad.XYZCoordinate{X: -1},
	RotateAngle:  90,
	// towards front half height
	TranslateFactor: scad.XYZCoordinate{Y: -0.5},
}

var SideTranslationBack = SideTranslation{
	// around X axis, the other direction
	RotateVector: scad.XYZCoordinate{X: 1},
	RotateAngle:  90,
	// towards back half height
	TranslateFactor: scad.XYZCoordinate{Y: 0.5},
}

var SideTranslationByCount = []SideTranslation{
	0: {},
	1: SideTranslationTop,
	2: SideTranslationBack,
	3: SideTranslationLeft,
	4: SideTranslationRight,
	5: SideTranslationFront,
	6: SideTranslationBottom,
}

type SideDimples struct {
	Name       string
	SideLength float64
	Dimple     Dimple
}

var ExampleSideDimples = SideDimples{
	Name:       "example_side_dimples",
	SideLength: 50,
	Dimple: Dimple{
		Diameter: 10,
		Depth:    8,
	},
}

func (s SideDimples) SCADWriter() scad.SCADWriter {
	sideDimpleObjects := make([]scad.Module, 6)
	for sideNumber, _ := range sideDimpleObjects {
		// there is no 0-dimple side
		dimpleCount := sideNumber + 1
		dimples := Dimples{
			Dimple: Dimple{
				Diameter: s.Dimple.Diameter,
				Depth:    s.Dimple.Depth,
			},
			Length:          s.SideLength,
			DimplePositions: DimplePositionsForCount[dimpleCount],
		}

		sideTranslation := SideTranslationByCount[dimpleCount]
		rotated := scad.RotationOf(
			float64(sideTranslation.RotateAngle),
			sideTranslation.RotateVector.X,
			sideTranslation.RotateVector.Y,
			sideTranslation.RotateVector.Z,
			dimples,
		)
		translated := scad.TranslationWithName(
			fmt.Sprintf("dimples_side_%d", sideNumber),
			s.SideLength*sideTranslation.TranslateFactor.X,
			s.SideLength*sideTranslation.TranslateFactor.Y,
			s.SideLength*sideTranslation.TranslateFactor.Z,
			rotated,
		)

		sideDimpleObjects[sideNumber] = translated
	}

	return scad.Union{
		Name: scad.Name{
			Specified: s.Name,
			Default:   "side_dimples",
		}.String(),
		Children: sideDimpleObjects,
	}.SCADWriter()
}
