package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Первые 500 простых чисел limit = 3572
// Первые 10000 простых чисел limit = 104730
// Первые 100000 простых чисел limit = 1299720

func main() {
	var n string
	if _, err := fmt.Fscan(os.Stdin, &n); err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	var limit int
	var err error
	if limit, err = strconv.Atoi(n); err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	nNumbers := natural_number1(limit)
	fmt.Println("")
	fmt.Println("Count alg1: ", len(nNumbers))

	nNumbers = natural_number2(limit)
	fmt.Println("")
	fmt.Println("Count alg2: ", len(nNumbers))

}

//goos: windows
//goarch: amd64
//pkg: gb_go
//BenchmarkNatural1
//BenchmarkNatural1-8   	       1	41510011100 ns/op
//PASS
func natural_number1(limit int) []int {
	var nNumbers []int

	nNumbers = append(nNumbers, 2, 3, 5)
	//fmt.Printf("%v", nNumbers)
iter:
	for i := 6; i <= limit; i++ {
		if !(i%2 == 0 || i%3 == 0 || i%5 == 0) {
			// Сложность перебора O(N)
			for _, nn := range nNumbers {
				if i%nn == 0 {
					continue iter
				}
			}
			nNumbers = append(nNumbers, i)
			//fmt.Printf(" %d", i)
		}
	}
	return nNumbers
}

//goos: windows
//goarch: amd64
//pkg: gb_go
//BenchmarkNatural2
//BenchmarkNatural2-8   	       6	 170876233 ns/op
//PASS
// Алгоритм значительно быстрее. Не имеет смысла перебирать натуральные числа как делитель более корня от i
func natural_number2(limit int) []int {
	var nNumbers []int

	nNumbers = append(nNumbers, 2, 3, 5)
	//fmt.Printf("%v", nNumbers)
iter:
	for i := 6; i <= limit; i++ {
		if !(i%2 == 0 || i%3 == 0 || i%5 == 0) {

			for _, nn := range nNumbers {
				//fmt.Println(int(math.Floor(math.Sqrt(float64(i)))))
				if i%nn == 0 {
					continue iter
				}
				// Если мы тут значит числом может быть простым
				// Но нет смысла искать делитетль больше чем корень от числа.
				if nn > int(math.Floor(math.Sqrt(float64(i)))) {
					break
				}
			}
			nNumbers = append(nNumbers, i)
			//fmt.Printf(" %d", i)
		}
	}
	return nNumbers
}
