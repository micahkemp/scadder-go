package component

import "github.com/micahkemp/scadder-go/renderer"

type RenderFields interface {
	Fields() renderer.Fields
	Render(path string, f renderer.Fields) (ok bool)
}

// this only exists so we can have effectively define methods for the RenderFields interface
type Component struct {
	RenderFields
}

func RenderFieldsComponent(r RenderFields) Component {
	return Component{r}
}
