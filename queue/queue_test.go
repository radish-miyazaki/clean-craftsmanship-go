package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		queue    *Queue
		expected bool
	}{
		{
			name:     "empty queue will return true",
			queue:    NewQueue([]int{}),
			expected: true,
		},
		{
			name:     "queue with one element will return false",
			queue:    NewQueue([]int{0}),
			expected: false,
		},
		{
			name:     "queue with two elements will return false",
			queue:    NewQueue([]int{0, 0}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.queue.IsEmpty())
		})
	}
}

func TestGetSize(t *testing.T) {
	tests := []struct {
		name     string
		queue    *Queue
		expected int
	}{
		{
			name:     "empty queue will return 0",
			expected: 0,
			queue:    NewQueue([]int{}),
		},
		{
			name:     "queue with one element will return 1",
			expected: 1,
			queue:    NewQueue([]int{0}),
		},
		{
			name:     "queue with two elements will return 2",
			expected: 2,
			queue:    NewQueue([]int{0, 0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.queue.GetSize())
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name     string
		queue    *Queue
		expected int
		count    int
	}{
		{
			name:     "push one element to empty queue, GetSize will return 1",
			expected: 1,
			queue:    NewQueue([]int{}),
			count:    1,
		},
		{
			name:     "push twice to empty queue, GetSize will return 2",
			expected: 2,
			queue:    NewQueue([]int{}),
			count:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.count; i++ {
				tt.queue.Push(0)
			}

			assert.Equal(t, tt.expected, tt.queue.GetSize())
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		queue    *Queue
		error    error
		expected []int
		count    int
	}{
		{
			name:  "pop to empty queue, will return UnderFlowErr",
			queue: NewQueue([]int{}),
			count: 1,
			error: &UnderFlowErr{},
		},
		{
			name:     "pop once to queue with 99, will return 99",
			queue:    NewQueue([]int{99}),
			count:    1,
			expected: []int{99},
		},
		{
			name:     "pop twice to queue with 99 88, will return 99 88 (FIFO)",
			queue:    NewQueue([]int{99, 88}),
			count:    2,
			expected: []int{99, 88},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var results []int

			for i := 0; i < tt.count; i++ {
				val, err := tt.queue.Pop()
				if tt.error != nil {
					assert.Equal(t, tt.error, err)
					return
				}

				results = append(results, val)
			}

			assert.Equal(t, tt.expected, results)
		})
	}
}
