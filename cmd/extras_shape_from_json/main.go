package main

import (
	"github.com/micahkemp/scad/scad"
	"github.com/micahkemp/scad/scad/extras"
	"os"
)

func main() {
	configType := os.Args[1]
	configFile := os.Args[2]
	outputName := os.Args[3]

	object := scad.NewOfTypeFromJSONFile(extras.ExampleConfigs, configType, outputName, configFile)
	object.Render(outputName)
}
