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
	ll.Remove(ll.Front())
	// [2]

	fmt.Fprintf(os.Stderr, "DEBUGPRINT[1]: queue.go:10: ll=%+v\n", ll.Front())
}
