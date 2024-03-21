package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonWhereNameToBeValueGiven(t *testing.T) {
	name := "Aaditya"
	person := NewPerson(name)
	assert.Equal(t, person.GetName(), name)
}
