package bug

import (
	"strings"
)

type Person struct {
	ID    int64
	Name  string
	Email *string
}

func (s Person) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + Inspect(s.ID, true)
	res[1] = "Name: " + Inspect(s.Name, true)
	res[2] = "Email: " + Inspect(s.Email, true)
	return strings.Join(res, ", ")
}
