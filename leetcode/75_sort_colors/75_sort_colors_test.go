package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type Test struct {
	name  string
	input []int
}

type Solution struct {
	name string
	fn   func([]int)
}

func TestSortColors(t *testing.T) {
	for _, sol := range []Solution{
		{"sortColorsMap", sortColorsMap},
		{"sortColorsOnePass", sortColorsOnePass},
	} {
		for _, tt := range []Test{
			{
				name:  "default case",
				input: []int{2, 0, 2, 1, 1, 0},
			},
			{
				name:  "sorted case",
				input: []int{0, 0, 1, 1, 2, 2, 2},
			},
			{
				name:  "empty case",
				input: []int{},
			},
			{
				name:  "reversed case",
				input: []int{2, 2, 2, 1, 1, 0, 0, 0},
			},
			{
				name:  "only 0 and 2 case",
				input: []int{2, 0, 2, 0, 0, 2},
			},
		} {
			t.Run(fmt.Sprintf("%s  %s", sol.name, tt.name), func(t *testing.T) {
				got := make([]int, len(tt.input))
				want := make([]int, len(tt.input))
				copy(got, tt.input)
				copy(want, tt.input)
				sol.fn(got)
				sort.Ints(want)

				if !reflect.DeepEqual(want, got) {
					t.Errorf("\nWanted:\t%+v\nGot:\t%+v", want, got)
				}
			})
		}
	}
}
