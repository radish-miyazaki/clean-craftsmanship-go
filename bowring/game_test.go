package bowring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func rollMany(g *Game, n int, pins int) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}

func rollSpare(g *Game) {
	rollMany(g, 2, 5)
}

func rollStrike(g *Game) {
	g.Roll(10)
}

func TestScore(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		setup    func(*Game)
	}{
		{
			name:     "Threw 20 times and knocked down 0 pins.",
			expected: 0,
			setup: func(g *Game) {
				rollMany(g, 20, 0)
			},
		},
		{
			name:     "Threw 20 times and knocked down 20 pins.",
			expected: 20,
			setup: func(g *Game) {
				rollMany(g, 20, 1)
			},
		},
		{
			name:     "Threw 20 times and knocked down 7 pins after one spare, and the rest gutter",
			expected: 24,
			setup: func(g *Game) {
				rollSpare(g)
				g.Roll(7)
				rollMany(g, 17, 0)
			},
		},
		{
			name:     "Threw 20 times and knocked down 2, 3 pins after one strike, and the rest gutter",
			expected: 20,
			setup: func(g *Game) {
				rollStrike(g)
				g.Roll(2)
				g.Roll(3)
				rollMany(g, 16, 0)
			},
		},
		{
			name:     "perfect game",
			expected: 300,
			setup: func(g *Game) {
				rollMany(g, 12, 10)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame()
			tt.setup(g)
			assert.Equal(t, g.Score(), tt.expected)
		})
	}
}
