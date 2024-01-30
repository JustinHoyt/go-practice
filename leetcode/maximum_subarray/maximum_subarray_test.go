package main

import (
	"testing"
)

type Test struct {
	name string
	want int
	input []int
}

func TestValidateBST(t *testing.T) {
	tests := []Test{
		{
			name: "empty array case",
			input: []int{},
			want: 0,
		},
		{
			name: "all negatives case",
			input: []int{-2,-3,-1},
			want: -1,
		},
		{
			name: "negative elements in a positive window slice case",
			input: []int{-2,1,-3,4,-1,2,1,-5,4},
			want: 6,
		},
	}

	for _, tt := range tests {
		got := maxSubArray1(tt.input)

		if got != tt.want {
			t.Errorf("%s failed\nwant maxSubArray(%+v) = %+v\ngot = %+v", tt.name, tt.input, tt.want, got)
		}
	}
}
