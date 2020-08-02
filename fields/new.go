package fields

type Fields struct {
	values map[string]string
}

func New(v map[string]string) Fields {
	return Fields{v}
}
