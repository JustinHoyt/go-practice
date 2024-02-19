package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestSortAnArray(t *testing.T) {
	for _, sol := range []struct {
		name string
		fn   func([]int)
	}{
		{"moveZeroes", moveZeroes},
		{"moveZeroesIter", moveZeroesIter},
	} {
		for _, tt := range []struct {
			name  string
			input []int
			want  []int
		}{
			{
				name:  "default",
				input: []int{0, 1, 0, 3, 12},
				want:  []int{1, 3, 12, 0, 0},
			},
			{
				name:  "long",
				input: []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0},
				want:  []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0},
			},
			{
				name:  "already done",
				input: []int{1, 0, 0},
				want:  []int{1, 0, 0},
			},
		} {
			t.Run(fmt.Sprintf("%s  %s", sol.name, tt.name), func(t *testing.T) {
				got := slices.Clone(tt.input)
				sol.fn(got)

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("\ninput:\t%+v\ngot:\t%+v\nwant\t%+v", tt.input, got, tt.want)
				}
			})
		}
	}
}
