package scad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ExampleToJSON(examples ExampleTypes) {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <type> <output file>\n", os.Args[0])
	}

	componentType := os.Args[1]
	filename := os.Args[2]

	example := examples[componentType]

	exampleJSON, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshal error: %s", err)
	}

	// OpenFile used specifically to enforce not overwriting an existing file, but instead failing
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		log.Fatalf("(ExampleToJSON) %s: %s\n", filename, err)
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", exampleJSON); err != nil {
		log.Fatalf("(ExampleToJSON) %s: %s\n", filename, err)
	}
}

func RenderFromJSONT(examples ExampleTypes) {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <type> <input file>\n", os.Args[0])
	}

	componentType := os.Args[1]
	jsonFilename := os.Args[2]

	jsonFileContents, err := ioutil.ReadFile(jsonFilename)
	if err != nil {
		log.Fatalf("input file (%s) error: %s", jsonFilename, err)
	}

	component, ok := examples[componentType]
	if !ok {
		log.Fatalf("type %s not known", componentType)
	}

	dec := json.NewDecoder(bytes.NewReader(jsonFileContents))
	dec.DisallowUnknownFields()

	err = dec.Decode(&component)
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}

	component.Render(".")
}
