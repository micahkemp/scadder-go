package scad

// seenTracker keeps track of strings that have been seen.
type seenTracker map[string]bool

// trackSeen returns true if the string has already been seen.
func (s seenTracker) trackSeen(input string) (seen bool) {
	if s[input] {
		return true
	}

	s[input] = true
	return false
}
