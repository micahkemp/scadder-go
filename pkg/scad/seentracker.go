package scad

// seenTracker keeps track of strings that have been seen.
type seenTracker map[string]bool

func (s seenTracker) trackPath(input string) (seen bool) {
	if s[input] {
		return false
	}

	s[input] = true
	return true
}
