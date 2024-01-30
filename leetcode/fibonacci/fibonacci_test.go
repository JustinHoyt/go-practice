package main

import "testing"

type testCase = struct {
	input int
	expected int
}

func runTests(t *testing.T, tests []testCase) {
	for _, test := range tests {
		if fibonacci(test.input) != test.expected {
			t.Errorf("%d expected to be %d", test.input, test.expected)
		}
	}
}

func TestFirstNumbers(t * testing.T) {
	runTests(t, []testCase{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	})
}
