package scadmodule

type component struct {
	name string
}

func (c component) Name() string {
	return c.Name()
}

func (c component) template() string {
	return "test"
}
