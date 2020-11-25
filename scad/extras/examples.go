package extras

import "github.com/micahkemp/scad/scad"

var ExampleConfigs = map[string]scad.ConfigConstructor{
	"tube": TubeConfigExample,
}
