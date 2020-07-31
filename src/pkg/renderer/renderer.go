package renderer

type Renderer struct {
	Template string
	Name     string
}

type Fields map[string]string

// enforces that both Template and Name are non-empty strings.
// while it's possible to create an instance of Renderer with
// such values, it's not valid to template them as such.
// it is up to the creating logic to ensure they are valid at creation time.
func (r Renderer) Render(path string, f Fields) (ok bool) {
	if r.Template == "" || r.Name == "" {
		return false
	}

	return true
}
