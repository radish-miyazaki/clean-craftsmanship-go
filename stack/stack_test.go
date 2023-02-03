package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		stack    Stack
		count    int
		expected []int
		error    error
	}{
		{
			name:     "pop stack with only 99, will return 99",
			stack:    *NewStack([]int{99}),
			count:    1,
			expected: []int{99},
		},
		{
			name:     "pop stack with only 88 will return 88",
			stack:    *NewStack([]int{88}),
			count:    1,
			expected: []int{88},
		},
		{
			name:     "pop stack with 99 and 88, will return 88 and 99 (FILO)",
			stack:    *NewStack([]int{99, 88}),
			count:    2,
			expected: []int{88, 99},
		},
		{
			name:     "pop empty stack, will return StackUnderFlowErr",
			stack:    *NewStack([]int{}),
			expected: []int{},
			count:    1,
			error:    &UnderFlowErr{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var results []int
			for i := 0; i < tt.count; i++ {
				val, err := tt.stack.Pop()
				if err != nil {
					assert.Equal(t, err, tt.error)
					return
				}

				results = append(results, val)
			}

			assert.Equal(t, tt.expected, results)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		stack    Stack
	}{
		{
			name:     "stack with one element will return false",
			expected: false,
			stack:    *NewStack([]int{0}),
		},
		{
			name:     "stack without element, will return true",
			expected: true,
			stack:    *NewStack([]int{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.stack.IsEmpty())
		})
	}
}

func TestGetSize(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		stack    Stack
	}{
		{
			name:     "stack with one element will return 1",
			expected: 1,
			stack:    *NewStack([]int{0}),
		},
		{
			name:     "stack with two elements will return 2",
			expected: 2,
			stack:    *NewStack([]int{0, 0}),
		},
		{
			name:     "stack without element will return 0",
			expected: 0,
			stack:    *NewStack([]int{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.stack.GetSize())
		})
	}
}
