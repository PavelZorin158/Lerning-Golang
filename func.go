package main

import "fmt"

func sumInt(a ...int) (int, int) {
	// передача переменного кол-ва параметров создает срез []int
	// возвращает кол-во параметров и их сумму
	var sum int
	for _, v := range a {
		sum += v
	}
	return len(a), sum
}
func main() {
	fmt.Println(sumInt(1, 2, 3))
}
