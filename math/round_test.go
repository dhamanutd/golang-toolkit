package math_test

import (
	"testing"

	"github.com/dhamanutd/golang-toolkit/math"
	"github.com/magiconair/properties/assert"
)

func TestRound(t *testing.T) {
	round := math.Round(777, -2)
	assert.Equal(t, float64(800), round)
}
