package sine

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const EPSILON = 0.000000001

type SineTaylorCalculator struct{}

func TestSine(t *testing.T) {
	tests := []struct {
		expected float64
		arg      float64
	}{
		{
			expected: 0.0,
			arg:      0.0,
		},
		{
			expected: 0.0,
			arg:      math.Pi,
		},
		{
			expected: 1.0,
			arg:      math.Pi / 2,
		},
		{
			expected: 0.8660254038,
			arg:      math.Pi / 3,
		},
		{
			expected: 0.7071067812,
			arg:      math.Pi / 4,
		},
		{
			expected: 0.5877852533,
			arg:      math.Pi / 5,
		},
		{
			expected: -1.0,
			arg:      3 * math.Pi / 2,
		},
		{
			expected: 0.0,
			arg:      2 * math.Pi,
		},
		{
			expected: 0.0,
			arg:      3 * math.Pi,
		},
		{
			expected: 0.0,
			arg:      -math.Pi,
		},
		{
			expected: -1.0,
			arg:      -math.Pi / 2,
		},
		{
			expected: 1.0,
			arg:      -3 * math.Pi / 2,
		},
		{
			expected: 0.0,
			arg:      -1000 * math.Pi,
		},
		{
			expected: 0.8660254038,
			arg:      1000*math.Pi + math.Pi/3,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			checkSin(t, tt.arg, tt.expected)
		})
	}
}

//func TestTaylorTerms(t *testing.T) {
//	spy := NewSineTaylorCalcSpy()
//	r := Math.Random() * math.Pi
//	for n := 1; n <= 10; n++ {
//		assert.Equal(t, n - 1, spy.GetSignPower()
//		assert.Equal(t, EPSILON, spy.GetR())
//		assert.Equal(t, 2 * n - 1, c.GetRPower())
//		assert.Equal(t, 2 * n - 1, c.GetFac())
//	}
//}

func checkSin(t *testing.T, rad float64, sin float64) {
	assert.InDelta(t, sin, Sin(rad), EPSILON)
}
