package main

import (
	"testing"
)

type Test struct {
	name string
	want int
	gas  []int
	cost []int
}

func TestValidateBST(t *testing.T) {
	tests := []Test{
		{
			name: "default case",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "invalid case",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			name: "big case",
			gas:  LongGas,
			cost: LongCost,
			want: 6690,
		},
		{
			name: "zeros case",
			gas:  []int{2, 0, 0, 0, 0},
			cost: []int{0, 1, 0, 0, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := canCompleteCircuit(tt.gas, tt.cost)

		if got != tt.want {
			t.Errorf("%s failed\nwant canCompleteCircuit(%+v, %+v) = %+v\ngot = %+v", tt.name, tt.gas, tt.cost, tt.want, got)
		}
	}
}
