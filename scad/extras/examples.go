package extras

import "github.com/micahkemp/scad/scad"

var ExampleConfigs = map[string]scad.ConfigConstructor{
	"nesting_tubes": NestingTubesConfigExample,
	"tube":          TubeConfigExample,
}
