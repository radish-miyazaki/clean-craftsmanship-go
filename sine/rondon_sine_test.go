package sine

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTaylorTerms(t *testing.T) {
	spy := NewTaylorCalculatorSpy(TaylorCalculator{})

	for n := 1; n <= 10; n++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Float64() * math.Pi

		spy.CalculateTerm(r, n)

		assert.Equal(t, n-1, spy.GetSignPower())
		assert.Equal(t, r, spy.GetR())
		assert.Equal(t, 2*n-1, spy.GetRPower())
		assert.Equal(t, 2*n-1, spy.GetFac())
	}
}

func TestSineInRange(t *testing.T) {
	c := TaylorCalculator{}
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		r := math.Pi*4 - (2 * math.Pi)

		sinr := c.Sin(r)

		assert.Less(t, sinr, 1.0)
		assert.Greater(t, sinr, -1.0)
	}
}

func TestPythagoreanIdentity(t *testing.T) {
	c := TaylorCalculator{}
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		r := math.Pi*4 - (2 * math.Pi)

		sinr := c.Sin(r)
		cosr := c.Cos(r)
		assert.InDelta(t, sinr*sinr+cosr*cosr, 1.0, 0.0001)
	}
}
