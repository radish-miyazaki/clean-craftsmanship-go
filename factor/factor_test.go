package factor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorsOf(t *testing.T) {
	tests := []struct {
		name     string
		expected []int
		arg      int
	}{
		{
			name:     "with 1, return will nil",
			expected: []int(nil),
			arg:      1,
		},
		{
			name:     "with 2, return will [2]",
			expected: []int{2},
			arg:      2,
		},
		{
			name:     "with 3, return will [3]",
			expected: []int{3},
			arg:      3,
		},
		{
			name:     "with 4, return will [2, 2]",
			expected: []int{2, 2},
			arg:      4,
		},
		{
			name:     "with 5, return will [5]",
			expected: []int{5},
			arg:      5,
		},
		{
			name:     "with 6, return will [2, 3]",
			expected: []int{2, 3},
			arg:      6,
		},
		{
			name:     "with 7, return will [7]",
			expected: []int{7},
			arg:      7,
		},
		{
			name:     "with 8, return will [2, 2, 2]",
			expected: []int{2, 2, 2},
			arg:      8,
		},
		{
			name:     "with 9, return will [3, 3]",
			expected: []int{3, 3},
			arg:      9,
		},
		{
			name:     "with 133, return will [7, 19]",
			expected: []int{7, 19},
			arg:      133,
		},
		{
			name:     "with 35281, return will [35281]",
			expected: []int{35281},
			arg:      35281,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, FactorsOf(tt.arg), tt.expected)
		})
	}
}
