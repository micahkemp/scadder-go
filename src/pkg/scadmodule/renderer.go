package scadmodule

type Renderer interface {
	Render(outputPath string) bool
}
