package fib

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		expected big.Int
		arg      int
	}{
		{
			name:     "received 0",
			arg:      0,
			expected: *big.NewInt(1),
		},
		{
			name:     "received 1",
			arg:      1,
			expected: *big.NewInt(1),
		},
		{
			name:     "received 2",
			arg:      2,
			expected: *big.NewInt(2),
		},
		{
			name:     "received 3",
			arg:      3,
			expected: *big.NewInt(3),
		},
		{
			name:     "received 10",
			arg:      9,
			expected: *big.NewInt(55),
		},
		{
			name:     "received 40",
			arg:      40,
			expected: *big.NewInt(165580141),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Fib(tt.arg))
		})
	}
}
