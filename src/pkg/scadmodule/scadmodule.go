package scadmodule

import (
	"fmt"
	"path"
	"regexp"
)

type SCADModule struct {
	Name string
}

// the match for module name validity is performed when asked for its filename
// the module name's validity doesn't even mean anything until it's attempted
// to be written disk
func (s SCADModule) filename() (filename string, ok bool) {
	ok = regexp.MustCompile("^[a-z][a-z0-9_]*$").MatchString(s.Name)
	filename = fmt.Sprintf("%s.scad", s.Name)
	return
}

func (s SCADModule) filepath(outputPath string, filename string) string {
	return path.Join(outputPath, filename)
}

func (s SCADModule) Render(outputPath string) bool {
	// dummy use of outputPath to avoid warnings
	if outputPath == "" {
		return false
	}

	return true
}
