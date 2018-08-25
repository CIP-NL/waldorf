package waldorf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObservation_IsFalse(t *testing.T) {
	waldorf := Observe()
	assert.True(t, waldorf.IsFalse(1 == 2, "because math"))
	assert.False(t, waldorf.IsFalse(1 == 1, "now match is good"))
}

func TestObservation_IsTrue(t *testing.T) {
	waldorf := Observe()
	assert.True(t, waldorf.IsTrue(1 == 1, "gud math"))
	assert.False(t, waldorf.IsTrue(1 == 2, "bad math"))
}

func TestObservation_Ridicule(t *testing.T) {
	waldorf := Observe()

	const (
		notRidiculed = "Are numbers equal?"
		incorrectOne = "are Jesus slippers Cool?"
		incorrectTwo = "Switserland is next to Sweden right?"
	)

	// correct observation, should not be in the list
	waldorf.IsTrue(1 == 1, notRidiculed)

	// incorrect observations, should be returned
	waldorf.IsTrue(1 == 2, incorrectOne)
	waldorf.IsFalse("A" == "A", incorrectTwo)

	everythingWrong := waldorf.Ridicule()
	assert.NotContains(t, everythingWrong, notRidiculed)
	assert.Contains(t, everythingWrong, incorrectOne)
	assert.Contains(t, everythingWrong, incorrectTwo)
}
