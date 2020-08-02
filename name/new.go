package name

type Name string

func New(s string) (n Name, ok bool) {
	if valid(s) {
		n = Name(s)
		ok = true
	}

	return
}

func FirstValid(strings ...string) (n Name, ok bool) {
	for _, s := range strings {
		n, ok = New(s)
		if ok {
			return
		}
	}

	return
}
