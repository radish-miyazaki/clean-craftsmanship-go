package wordwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		arg0     string
		arg1     int
	}{
		{
			name:     "pass empty string",
			expected: "",
			arg0:     "",
			arg1:     1,
		},
		{
			name:     "pass one char with chars when per line one",
			expected: "x",
			arg0:     "x",
			arg1:     1,
		},
		{
			name:     "pass two chars when per line one",
			expected: "x\nx",
			arg0:     "xx",
			arg1:     1,
		},
		{
			name:     "pass two chars when per line two",
			expected: "xx",
			arg0:     "xx",
			arg1:     2,
		},
		{
			name:     "pass three chars when per line one",
			expected: "x\nx\nx",
			arg0:     "xxx",
			arg1:     1,
		},
		{
			name:     "pass three chars when per line two",
			expected: "xx\nx",
			arg0:     "xxx",
			arg1:     2,
		},
		{
			name:     "pass three chars when per line three",
			expected: "xxx",
			arg0:     "xxx",
			arg1:     3,
		},
		{
			name:     "pass two chars include space when per line one",
			expected: "x\nx",
			arg0:     "x x",
			arg1:     1,
		},
		{
			name:     "pass two chars include space when per line two",
			expected: "x\nx",
			arg0:     "x x",
			arg1:     2,
		},
		{
			name:     "pass three chars include space when per line one",
			expected: "x\nx\nx",
			arg0:     "x x x",
			arg1:     1,
		},
		{
			name:     "pass three chars include space when per line two",
			expected: "x\nx\nx",
			arg0:     "x x x",
			arg1:     2,
		},
		{
			name:     "pass three chars include space when per line three",
			expected: "x x\nx",
			arg0:     "x x x",
			arg1:     3,
		},
		{
			name:     "pass three chars include space when per line four",
			expected: "x x\nx",
			arg0:     "x x x",
			arg1:     4,
		},
		{
			name:     "pass three chars include space when per line five",
			expected: "x x x",
			arg0:     "x x x",
			arg1:     5,
		},
		{
			name:     "pass four chars include space when per line one",
			expected: "x\nx\nx\nx",
			arg0:     "xx xx",
			arg1:     1,
		},
		{
			name:     "pass four chars include space when per line two",
			expected: "xx\nxx",
			arg0:     "xx xx",
			arg1:     2,
		},
		{
			name:     "pass four chars include space when per line three",
			expected: "xx\nxx",
			arg0:     "xx xx",
			arg1:     3,
		},
		{
			name:     "pass four chars include space when per line four",
			expected: "xx\nxx",
			arg0:     "xx xx",
			arg1:     4,
		},
		{
			name:     "pass four chars include space when per line five",
			expected: "xx xx",
			arg0:     "xx xx",
			arg1:     5,
		},
		{
			name:     "pass six chars include space when per line one",
			expected: "x\nx\nx\nx\nx\nx",
			arg0:     "xx xx xx",
			arg1:     1,
		},
		{
			name:     "pass six chars include space when per line two",
			expected: "xx\nxx\nxx",
			arg0:     "xx xx xx",
			arg1:     2,
		},
		{
			name:     "pass six chars include space when per line three",
			expected: "xx\nxx\nxx",
			arg0:     "xx xx xx",
			arg1:     3,
		},
		{
			name:     "pass six chars include space when per line four",
			expected: "xx\nxx\nxx",
			arg0:     "xx xx xx",
			arg1:     4,
		},
		{
			name:     "pass six chars include space when per line five",
			expected: "xx xx\nxx",
			arg0:     "xx xx xx",
			arg1:     5,
		},
		{
			name:     "pass six chars include space when per line six",
			expected: "xx xx\nxx",
			arg0:     "xx xx xx",
			arg1:     6,
		},
		{
			name:     "pass six chars include space when per line seven",
			expected: "xx xx\nxx",
			arg0:     "xx xx xx",
			arg1:     7,
		},
		{
			name:     "pass six chars include space when per line eight",
			expected: "xx xx xx",
			arg0:     "xx xx xx",
			arg1:     8,
		},
		{
			name:     "pass string more than to chars per line",
			expected: "Four\nscore\nand\nseven\nyears\nago our",
			arg0:     "Four score and seven years ago our",
			arg1:     7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Wrap(tt.arg0, tt.arg1))
		})
	}
}
