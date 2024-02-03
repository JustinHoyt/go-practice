package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSortAnArray(t *testing.T) {
	for _, sol := range []struct {
		name string
		fn   func([]int) []int
	}{
		{"merge sort", mergeSort},
	} {
		for _, tt := range []struct {
			name  string
			input []int
		}{
			{"positive nums", []int{2, 4, 9, 8, 7, 1, 9, 4, 6}},
			{"empty list", []int{}},
			{"negative nums", []int{5, 2, -100}},
			{"duplicate nums", []int{12, 9, 48, 48, 48, 7, 6, 89}},
		} {
			t.Run(fmt.Sprintf("%s  %s", sol.name, tt.name), func(t *testing.T) {
				got := make([]int, len(tt.input))
				want := make([]int, len(tt.input))
				copy(got, tt.input)
				copy(want, tt.input)
				got = sol.fn(tt.input)
				sort.Ints(want)

				if !reflect.DeepEqual(got, want) {
					t.Errorf("\ninput:\t%+v\ngot:\t%+v\nwant\t%+v", tt.input, got, want)
				}
			})
		}
	}
}
