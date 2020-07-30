package scadmodule

import (
	"fmt"
	"reflect"
	"strings"
)

type named interface {
	Name() string
}

func selectName(n named) string {
	if n.Name() != "" {
		return n.Name()
	}

	typeOfN := fmt.Sprintf("%s", reflect.TypeOf(n))
	typeComponents := strings.Split(typeOfN, ".")
	simpleTypeOfN := typeComponents[len(typeComponents)-1]

	return simpleTypeOfN
}
