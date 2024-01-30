package main

import (
	"fmt"
	"math"
)

func isPalindromeBrute(x int) bool {
	if x < 0 {
		return false
	}

	xStr := fmt.Sprint(x)

	for i, j := 0, len(xStr)-1; i < j; i, j = i+1, j-1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}

func isPalindromeIntMath(x int) bool {
	if x < 0 {
		return false
	}

	for x > 9 {
		length := int(math.Log10(float64(x))) + 1
		magnitude := int(math.Pow(10, float64(length)-1))
		var small_digit int = x % 10
		var big_digit int = x / magnitude
		if big_digit != small_digit {
			return false
		}
		// remove left most digit
		x -= magnitude * big_digit
		// remove right most digit
		x /= 10
	}
	return true
}

func main() {
	isPalindromeIntMath(1231)
}
