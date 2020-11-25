package scad

import (
	"io/ioutil"
	"log"
)

type ConfigConstructor interface {
	NewFromConfig(name string) DirPathRenderer
	NewFromJSON(n string, j []byte) DirPathRenderer
}

func ObjectTypeFromJSON(e map[string]ConfigConstructor, t string, n string, c []byte) DirPathRenderer {
	exampleConfig, ok := e[t]
	if !ok {
		log.Fatalf("Unknown object type: %s", t)
	}

	return exampleConfig.NewFromJSON(n, c)
}

func NewOfTypeFromJSONFile(e map[string]ConfigConstructor, t string, n string, j string) DirPathRenderer {
	configContents, err := ioutil.ReadFile(j)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	return ObjectTypeFromJSON(e, t, n, configContents)
}
