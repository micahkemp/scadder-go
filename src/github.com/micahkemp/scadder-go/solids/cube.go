package solids

import (
	"fmt"
	"github.com/micahkemp/scadder-go/renderer"
	"strconv"
)

type Cube struct {
	l, w, h float64
	renderer.Renderer
}

func (c Cube) Fields() renderer.Fields {
	return renderer.Fields{
		"size": fmt.Sprintf("[%s, %s, %s]",
			// use as few digits as possible, since these will often be int-like or low precision
			strconv.FormatFloat(c.l, 'f', -1, 64),
			strconv.FormatFloat(c.w, 'f', -1, 64),
			strconv.FormatFloat(c.h, 'f', -1, 64)),
	}
}
