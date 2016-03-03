package bug

import (
	"fmt"
)

func Inspect(arg interface{}, addType bool) string {
	var s string
	switch a := arg.(type) {
	case string:
		s = fmt.Sprintf("%#q", a)

	case *string:
		if a == nil {
			s = "<nil>"
		} else {
			s = Inspect(*a, false)
		}

	default:
		s = fmt.Sprintf("%v", a)
	}

	if addType {
		s += fmt.Sprintf(" (%T)", arg)
	}
	return s
}
