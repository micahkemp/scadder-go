package name

import "fmt"

const extension = "scad"

func (n Name) Filename() string {
	return fmt.Sprintf("%s.%s", n, extension)
}
