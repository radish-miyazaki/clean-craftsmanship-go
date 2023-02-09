package rental

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegularMovieOneDay(t *testing.T) {
	tests := []struct {
		name           string
		days           int
		expectedFee    int
		expectedPoints int
	}{
		{
			name:           "one day",
			days:           1,
			expectedFee:    150,
			expectedPoints: 1,
		},
		{
			name:           "two days",
			days:           2,
			expectedFee:    150,
			expectedPoints: 1,
		},
		{
			name:           "three days",
			days:           3,
			expectedFee:    150,
			expectedPoints: 1,
		},
		{
			name:           "four days",
			days:           4,
			expectedFee:    300,
			expectedPoints: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			c.AddRental("RegularMovie", tt.days)
			assert.Equal(t, tt.expectedFee, c.GetRentalFee())
			assert.Equal(t, tt.expectedPoints, c.GetRentalPoints())
		})
	}
}

func TestChildrenMovie(t *testing.T) {
	tests := []struct {
		name           string
		days           int
		expectedFee    int
		expectedPoints int
	}{
		{
			name:           "one day",
			days:           1,
			expectedFee:    100,
			expectedPoints: 1,
		},
		{
			name:           "four days",
			days:           4,
			expectedFee:    400,
			expectedPoints: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			c.AddRental("ChildrenMovie", tt.days)
			assert.Equal(t, tt.expectedFee, c.GetRentalFee())
			assert.Equal(t, tt.expectedPoints, c.GetRentalPoints())
		})
	}
}

func TestOneRegularAndOneChildrenMovie(t *testing.T) {
	tests := []struct {
		name           string
		regularDays    int
		childrenDays   int
		expectedFee    int
		expectedPoints int
	}{
		{
			name:           "both four days",
			regularDays:    4,
			childrenDays:   4,
			expectedFee:    700,
			expectedPoints: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			c.AddRental("RegularMovie", tt.regularDays)
			c.AddRental("ChildrenMovie", tt.childrenDays)
			assert.Equal(t, tt.expectedFee, c.GetRentalFee())
			assert.Equal(t, tt.expectedPoints, c.GetRentalPoints())
		})
	}
}
