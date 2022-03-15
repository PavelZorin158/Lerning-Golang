package main

import (
	"fmt"
)

func main() {
	var a map[string]int // так нельзя
	var b = map[string]int{"kat": 46, "Pav": 47}
	c := map[string]int{}
	d := make(map[string]int)
	c["vova"] = 158

	fmt.Println(a, "\n", b, "\n", c, "\n", d)

	//====================================================

	m := map[int]int{
		1: 10,
		3: 13,
		4: 14,
	}
	// ПРОВЕРКА НА НАЛИЧИЕ КЛЮЧА
	_, ok := m[1]
	if ok {
		fmt.Println(m[1]) // 10
	}

	if value, inMap := m[2]; inMap {
		fmt.Println(value) // Условие не выполняется
	}
	m[5] = 8     // добавляем элемент 8 с ключем 5
	delete(m, 3) // удаляем элемент с ключем 3

	// итерация по мапу
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(len(m), "элемента")

}
