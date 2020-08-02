package name

import "regexp"

// start with letter, end with letter
// digits and underscores permitted between
const validPattern = "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z]$"

func valid(s string) bool {
	return regexp.MustCompile(validPattern).MatchString(s)
}
