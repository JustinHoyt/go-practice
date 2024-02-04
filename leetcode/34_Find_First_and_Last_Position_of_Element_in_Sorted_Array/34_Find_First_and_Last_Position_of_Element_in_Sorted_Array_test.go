package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestSortAnArray(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input Input
		want  []int
	}{
		{
			name:  "Single match",
			input: Input{[]int{5, 7, 7, 8, 10}, 8},
			want:  []int{3, 3},
		},
		{
			name:  "no match",
			input: Input{[]int{5, 7, 7, 10}, 8},
			want:  []int{-1, -1},
		},
		{
			name:  "multiple matches",
			input: Input{[]int{5, 7, 7, 7, 10}, 7},
			want:  []int{1, 3},
		},
		{
			name:  "empty list",
			input: Input{[]int{}, 7},
			want:  []int{-1, -1},
		},
		{
			name:  "all match",
			input: Input{[]int{7, 7, 7, 7, 7, 7, 7, 7}, 7},
			want:  []int{0, 7},
		},
		{
			name:  "negative nums",
			input: Input{[]int{-15, -7, -7, 7, 10}, -7},
			want:  []int{1, 2},
		},
	} {
		t.Run(fmt.Sprintf("searchRange  %s", tt.name), func(t *testing.T) {
			got := searchRange(tt.input.nums, tt.input.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ninput:\t%+v\ngot:\t%+v\nwant\t%+v", tt.input, got, tt.want)
			}
		})
	}
}
