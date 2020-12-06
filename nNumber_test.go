package main

import "testing"

func BenchmarkNatural1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		_ = natural_number1(1299720)
	}
}

func BenchmarkNatural2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		natural_number2(1299720)
	}
}
