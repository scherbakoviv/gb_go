package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test := []int{5, 6, 3, 4, 8, 2, 0, 3, 45, 7, 8, 3, 4, 6, 3, 6, 78, 9, 1, 2, 4, 0, -1, -3, 4, -0}
	valid := []int{-3, -1, 0, 0, 0, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 6, 6, 6, 7, 8, 8, 9, 45, 78}

	sort := SortInsertClassic(test)
	fmt.Println(test)
	fmt.Println(sort)

	if intSliceEq(sort, valid) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")
	}
}

func GetRandomSlice() []int {
	a := []int{}
	for i := 0; i < 1000; i++ {
		a = append(a, rand.Intn(100))
	}
	return a
}

func SortBubble(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return []int{}
	}
	tmp := make([]int, len(a))
	copy(tmp, a)
	for i := 0; i < len(a); i++ {
		check := false
		for ii := i; ii < len(a); ii++ {
			if tmp[i] > tmp[ii] {
				tmp[i], tmp[ii] = tmp[ii], tmp[i]
				check = true
			}
		}
		if !check {
			break
		}
	}
	return tmp
}

func SortInsert(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return []int{}
	}
	tmp := make([]int, len(a))
	copy(tmp, a)

	for i := 1; i < len(tmp); i++ {
		x := 0
		//Поиск позиции, куда надо вставить элемент.
		for ii := i - 1; ii >= 0; ii-- {
			if tmp[ii] > tmp[i] {
				x++
			}else{
				//Нашли позицию, дальше проверять не обязательно.
				break
			}
		}
		if x > 0 {
			// Смещаем слайс на 1 и вставляем элемент.
			//t := tmp[i]
			copy(tmp[i-x+1:i+1], tmp[i-x:i])
			//tmp[i-x] = t
			tmp[i-x] = a[i] //есть же оригинал, зачем создавать t

		}
	}
	return tmp
}

func SortInsertT(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return []int{}
	}
	tmp := make([]int, len(a))
	copy(tmp, a)

	for i := 1; i < len(tmp); i++ {
		x := 0
		//Поиск позиции, куда надо вставить элемент.
		for ii := i - 1; ii >= 0; ii-- {
			if tmp[ii] > tmp[i] {
				x++
			}else{
				//Нашли позицию, дальше проверять не обязательно.
				break
			}
		}
		if x > 0 {
			// Смещаем слайс на 1 и вставляем элемент.
			t := tmp[i]
			copy(tmp[i-x+1:i+1], tmp[i-x:i])
			tmp[i-x] = t
		}
	}
	return tmp
}

func SortInsertClassic(a []int) []int {
	if a == nil {
		return nil
	}
	if len(a) == 0 {
		return []int{}
	}
	tmp := make([]int, len(a))
	copy(tmp, a)

	for i := 1; i < len(tmp); i++ {
		for ii := i - 1; ii >= 0; ii-- {
			if tmp[ii] > tmp[ii+1] {
				tmp[ii], tmp[ii+1] = tmp[ii+1], tmp[ii]
			}else{
				break
			}
		}
	}
	return tmp
}

func intSliceEq(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
