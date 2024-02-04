package main

import (
	"testing"
)

type Test struct {
	name  string
	want  int
	input []int
}

func TestValidateBST(t *testing.T) {
	tests := []Test{
		{
			name:  "default case",
			input: []int{100, 4, 200, 1, 3, 2},
			want:  4,
		},
		{
			name:  "duplicate case",
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:  9,
		},
		{
			name:  "empty case",
			input: []int{},
			want:  0,
		},
		{
			name:  "negative case",
			input: []int{-3, -100, -1, -2, -99},
			want:  3,
		},
	}

	for _, tt := range tests {
		got := longestConsecutive(tt.input)

		if got != tt.want {
			t.Errorf("%s failed\nwant longestConsecutive(%+v) = %+v\ngot = %+v", tt.name, tt.input, tt.want, got)
		}
	}
}
