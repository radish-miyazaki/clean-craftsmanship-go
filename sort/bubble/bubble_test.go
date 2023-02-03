package bubble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		arg      []int
		expected []int
	}{
		{
			name:     "after pass empty array, will return empty array",
			arg:      []int(nil),
			expected: []int(nil),
		},
		{
			name:     "after pass [1], will return [1]",
			arg:      []int{1},
			expected: []int{1},
		},
		{
			name:     "after pass [1, 2], will return [1, 2]",
			arg:      []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "after pass [2, 1], will return [1, 2]",
			arg:      []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "after pass [1, 2, 3], will return [1, 2, 3]",
			arg:      []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "after pass [2, 1, 3], will return [1, 2, 3]",
			arg:      []int{2, 1, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "after pass [3, 2, 1], will return [1, 2, 3]",
			arg:      []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "after pass [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3], will return [1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 7, 8, 9, 9, 9]",
			arg:      []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3},
			expected: []int{1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 7, 8, 9, 9, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Sort(tt.arg))
		})
	}
}
