package main

import "fmt"

func fa(ff func(int) int, c int) int {
	// принимает функцию fb и число 5
	// передает это число в принятую функцию и к полученному
	// из нее результату прибавляет 3
	return ff(c) + 3
}

func fb(a int) int {
	return a * 2
}

func main() {
	// передаем в фунцию fa 2 параметра: функцию fb и число 5
	fmt.Println(fa(fb, 5))

	// передаем в функцию fa функцию описываемую на месте
	res := fa(func(i int) int { return i * 3 }, 5)
	fmt.Println(res)

	// присваиваем переменной fc функцию
	fc := func(i int) int {
		i = i * 4
		return i
	}
	// передаем в функцию fa переменную fc
	fmt.Println(fa(fc, 5))
}
