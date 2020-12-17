package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSortBubble(b *testing.B) {
	rand.Seed(1000)
	s := make([]int, 1000)
	for n := 0; n < b.N; n++ {
		_ = FillRandomSlice(s)
		SortBubble(s)
	}
}

func BenchmarkSortInsert(b *testing.B) {
	rand.Seed(1000)
	s := make([]int, 1000)
	for n := 0; n < b.N; n++ {
		_ = FillRandomSlice(s)
		SortBubble(s)
	}
}

func BenchmarkSortInsertT(b *testing.B) {
	rand.Seed(1000)
	s := make([]int, 1000)
	for n := 0; n < b.N; n++ {
		_ = FillRandomSlice(s)
		SortBubble(s)
	}
}

func BenchmarkSortInsertClassic(b *testing.B) {
	rand.Seed(1000)
	s := make([]int, 1000)
	for n := 0; n < b.N; n++ {
		_ = FillRandomSlice(s)
		SortBubble(s)
	}
}

//BenchmarkSortBubble-8               1670            671214 ns/op           24570 B/op         12 allocs/op
//BenchmarkSortInsert-8               7500            187200 ns/op           24568 B/op         12 allocs/op
//BenchmarkSortInsertClassic-8        3247            368769 ns/op           24568 B/op         12 allocs/op

//BenchmarkSortBubble-8               1795            637883 ns/op           24570 B/op         12 allocs/op
//BenchmarkSortInsert-8               6332            187292 ns/op           24568 B/op         12 allocs/op
//BenchmarkSortInsertClassic-8        2930            375837 ns/op           24568 B/op         12 allocs/op
