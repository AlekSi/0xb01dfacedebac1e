package bug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrash(t *testing.T) {
	email := "elfrieda_abbott@example.org"
	person1 := Person{ID: 102, Name: "1", Email: &email}
	person2 := Person{ID: 102, Name: "2", Email: &email}
	assert.Equal(t, person1, person2) // crash here!
}
