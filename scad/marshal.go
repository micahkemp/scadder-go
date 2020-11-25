package scad

import (
	"bytes"
	"encoding/json"
	"log"
)

func UnmarshalStrict(j []byte, v interface{}) {
	dec := json.NewDecoder(bytes.NewReader(j))
	dec.DisallowUnknownFields()

	err := dec.Decode(&v)
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
}
