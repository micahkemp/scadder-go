package extras

import (
	"github.com/micahkemp/scad/scad"
	"log"
)

type TubeConfig struct {
	InnerDiameter float64
	OuterDiameter float64
	Height        float64
	Center        bool
}

var TubeConfigExample = TubeConfig{
	10,
	12,
	20,
	true,
}

type tube struct {
	TubeConfig
	scad.Difference
}

func NewTubeFromConfig(n string, t TubeConfig) tube {
	if t.InnerDiameter >= t.OuterDiameter {
		log.Fatalf("Tube inner diameter (%f) must be smaller than outer diameter (%f)",
			t.InnerDiameter, t.OuterDiameter)
	}

	outerCylinder := scad.NewCylinderByDiameter("outer", t.OuterDiameter, t.Height, t.Center)
	innerCylinder := scad.NewCylinderByDiameter("inner", t.InnerDiameter, t.Height, t.Center)

	return tube{
		t,
		outerCylinder.Subtract(n, innerCylinder),
	}
}

func (t TubeConfig) NewFromConfig(n string) scad.DirPathRenderer {
	return NewTubeFromConfig(n, t)
}

func (_ TubeConfig) NewFromJSON(n string, j []byte) scad.DirPathRenderer {
	var t TubeConfig
	scad.UnmarshalStrict(j, &t)
	return t.NewFromConfig(n)
}
