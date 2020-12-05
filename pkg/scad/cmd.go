package scad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CLIHandler(models Models) {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s exportJSON|renderJSON <model> <file>",
			os.Args[0],
		)
	}

	action := os.Args[1]
	modelName := os.Args[2]
	fileName := os.Args[3]

	model, ok := models[modelName]
	if !ok {
		log.Fatalf("Unknown model: %s", modelName)
	}

	switch action {
	case "exportJSON":
		exportJSON(model, fileName)
	case "renderJSON":
		renderJSON(model, fileName)
	default:
		log.Fatalf("Unknown action: %s", action)
	}
}

func exportJSON(model Module, filename string) {
	modelJSON, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshal error: %s", err)
	}

	// OpenFile used specifically to enforce not overwriting an existing file, but instead failing
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		log.Fatalf("(ExampleToJSON) %s: %s\n", filename, err)
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", modelJSON); err != nil {
		log.Fatalf("(ExampleToJSON) %s: %s\n", filename, err)
	}
}

func renderJSON(model Module, filename string) {
	modelJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("input file (%s) error: %s", filename, err)
	}

	dec := json.NewDecoder(bytes.NewReader(modelJSON))
	dec.DisallowUnknownFields()

	err = dec.Decode(&model)
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}

	model.SCADWriter().WriteSCAD("output")
}
