package extras

import "github.com/micahkemp/scad/scad"

type NestingTubesConfig struct {
	ExternalTube TubeConfig
	InternalTube TubeConfig
}

var NestingTubesConfigExample = NestingTubesConfig{
	ExternalTube: TubeConfig{
		InnerDiameter: 20,
		OuterDiameter: 25,
		Height:        100,
		Center:        true,
	},
	InternalTube: TubeConfig{
		InnerDiameter: 13,
		OuterDiameter: 18,
		Height:        100,
		Center:        true,
	},
}

type NestingTubes struct {
	NestingTubesConfig
	scad.Union
}

func NewNestingTubesFromConfig(n string, c NestingTubesConfig) NestingTubes {
	externalTube := NewTubeFromConfig("external", c.ExternalTube)
	internalTube := NewTubeFromConfig("internal", c.InternalTube)

	return NestingTubes{
		c,
		externalTube.Add(n, internalTube),
	}
}

func (c NestingTubesConfig) NewFromConfig(n string) scad.DirPathRenderer {
	return NewNestingTubesFromConfig(n, c)
}

func (_ NestingTubesConfig) NewFromJSON(n string, j []byte) scad.DirPathRenderer {
	var nt NestingTubesConfig
	scad.UnmarshalStrict(j, &nt)
	return nt.NewFromConfig(n)
}
