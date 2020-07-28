package renderer

import (
	"path"
)

type Renderer struct {
	OutputPath string
}

func (r *Renderer) filepath(filename string) string {
	return path.Join(r.OutputPath, filename)
}
