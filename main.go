package main

import (
	"fmt"
	"os"
	"strconv"
)
// Первые 500 простых чисел limit = 3572
// Первые 10000 простых чисел limit = 104730
// Первые 100000 простых чисел limit = 1299720
func main()  {
	var n string
	if _, err := fmt.Fscan(os.Stdin, &n); err != nil{
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	var limit int
	var err error
	if limit, err = strconv.Atoi(n); err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	var n_numbers []int

	n_numbers = append(n_numbers, 2,3,5)
	fmt.Printf("%v", n_numbers)
	iter:
	for i:=6; i <= limit; i++ {
		if !(i % 2 == 0 || i % 3 == 0 || i % 5 == 0) {
			for _, nn:= range n_numbers{
				if i % nn == 0 {
					continue iter
				}
			}
			n_numbers = append(n_numbers, i)
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println("")
	fmt.Println("Count: ", len(n_numbers))
}



