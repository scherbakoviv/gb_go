package main

import (
	"fmt"
	"os"
)

var fibos = make(map[int]int)
func main() {
	for {
		fmt.Println("Введите номер числа фибоначи:")
		var ns int
		if _, err := fmt.Fscan(os.Stdin, &ns); err != nil {
			panic(err)
		}
		fmt.Println(fibo(ns))
	}
}

func fibo(n int) int  {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if _, f := fibos[n]; f == false  {
		fibos[n] = fibo(n - 2) + fibo(n - 1)
	}

	return fibos[n]
}