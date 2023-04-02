package main

import "testing"

type testCase = struct {
	input    int
	expected bool
}

func runTests(t *testing.T, tests []testCase) {
	for _, palFn := range []func(int) bool{
		isPalindromeBrute,
		isPalindromeIntMath,
	} {
		for _, test := range tests {
			if palFn(test.input) != test.expected {
				t.Errorf("%d expected to be %t", test.input, test.expected)
			}
		}
	}
}

func TestOddLengthPalindrome(t *testing.T) {
	runTests(t, []testCase{
		{121, true},
		{122, false},
		{12121, true},
	})
}

func TestEvenLengthPalindrome(t *testing.T) {
	runTests(t, []testCase{
		{11, true},
		{12, false},
		{1221, true},
		{1231, false},
		{121121, true},
	})
}

func TestNegativePalindrome(t *testing.T) {
	runTests(t, []testCase{
		{-11, false},
	})
}
