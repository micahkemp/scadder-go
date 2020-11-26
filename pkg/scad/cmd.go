package scad

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ExampleToJSON(examples ExampleTypes) {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <type> <file>\n", os.Args[0])
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
