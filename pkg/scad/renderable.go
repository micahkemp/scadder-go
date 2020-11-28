package scad

type Renderer interface {
	Render(path string)
}

type renderable struct {
	renderer Renderer
}

func Renderable(r Renderer) renderable {
	return renderable{
		r,
	}
}
