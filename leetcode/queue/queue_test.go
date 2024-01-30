package main

import "testing"

func BenchmarkQueue(b *testing.B) {
    var q Queue
    for i := 0; i < b.N; i++ {
        q.Enqueue(i)
    }
    for i := 0; i < b.N; i++ {
        q.Dequeue()
    }
}

