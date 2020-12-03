package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main()  {
	var a string
	for{

		fmt.Println("Введите выражение, например 5 + 2:")
		if _, err := fmt.Fscan(os.Stdin, &a); err != nil {
			fmt.Println("Какая-то ошибка: %v", err)
			os.Exit(1)
		}
		r := regexp.MustCompile(`(?P<v1>-?\d+) ?(?P<operator>[+\-*\/^]) ?(?P<v2>-?\d+)`)
		m := r.FindStringSubmatch(a)
		if len(m) == 0 {
			fmt.Println("Только числа и операторы +,-,*,/ ")
			continue
		}
		v1 := m[r.SubexpIndex("v1")]
		v2 := m[r.SubexpIndex("v2")]
		operator := m[r.SubexpIndex("operator")]

		var var1, var2 int64
		var err error
		if var1, err = strconv.ParseInt(v1, 10, 64); err != nil {
			fmt.Println("Первый операнд должен быть числом")
			continue
		}

		if var2, err = strconv.ParseInt(v2, 10, 64); err != nil {
			fmt.Println("Второй операнд должен быть числом")
			continue
		}

		var result int64
		switch operator {
		case "+":
			result = var1 + var2
		case "-":
			result = var1 - var2
		case "*":
			result = var1 * var2
		case "/":
			if var2 == 0 {
				fmt.Println("Я бы с радостью, но хозяин не разрешает делить на ноль")
				continue
			}
			result = var1 / var2
		default:
			fmt.Println("Поддерживаются только операторы +,-,*,/")
			continue
		}

		fmt.Println("Равно: ", result)
	}
}

