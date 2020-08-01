package solids

import (
	"fmt"
	"github.com/micahkemp/scad/component"
	"github.com/micahkemp/scad/renderer"
	"strconv"
)

type cube struct {
	l, w, h float64
	renderer.Renderer
}

func New(l, w, h float64, name string) component.Component {
	if name == "" {
		name = "cube_component"
	}

	return component.Component{
		RenderFields: cube{
			l: l,
			w: w,
			h: h,
			Renderer: renderer.Renderer{
				Template: template,
				Name:     name,
			},
		},
	}
}

func (c cube) Fields() renderer.Fields {
	return renderer.Fields{
		"size": fmt.Sprintf("[%s, %s, %s]",
			// use as few digits as possible, since these will often be int-like or low precision
			strconv.FormatFloat(c.l, 'f', -1, 64),
			strconv.FormatFloat(c.w, 'f', -1, 64),
			strconv.FormatFloat(c.h, 'f', -1, 64)),
	}
}
