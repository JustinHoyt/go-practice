package main

import (
	"fmt"
)

// 1 2 3 4 5 6  7
// 1 1 2 3 5 8 13
func fibonacci(target int) int {
	memo := map[int]int{
		1: 1,
		2: 1,
	}

	var fibRec func(int) int
	fibRec = func(curr int) int {
		val, exists := memo[curr]
		if (exists) {
			return val
		}

		memo[curr] = fibRec(curr - 1) + fibRec(curr - 2)
		return memo[curr]
	}

	return fibRec(target)
}

func main() {
	fmt.Println(fibonacci(5))
}
