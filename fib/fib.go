package main

import (
	"main/tap"

	"github.com/samber/mo"
)

func fib(n int) mo.Result[int] {
	var memo = map[int]int{1: 1, 2: 1}
	var fib_rec func(int) int
	fib_rec = func(num_left int) int {
		if _, ok := memo[num_left]; !ok {
			memo[num_left] = fib_rec(num_left-1) + fib_rec(num_left-2)
		}
		return memo[num_left]
	}

	return mo.Ok(fib_rec(n))
}

func main() {
	fib(100).Map(tap.Tap(tap.UnaryPrintln[int]))
}
