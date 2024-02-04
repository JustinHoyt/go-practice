package main

import (
	"container/list"
	"fmt"
	"os"
)

func main() {
	ll := list.List{}

	ll.PushBack(1)
	// [1]
	ll.PushBack(2)
	// [1, 2]
	var val int = ll.Remove(ll.Front()).(int)
	// [2]
	fmt.Fprintf(os.Stderr, "DEBUGPRINT[7]: queue.go:16: val=%+v\n", val)
}
