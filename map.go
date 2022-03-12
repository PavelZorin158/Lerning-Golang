package main

import (
	"fmt"
)

func main() {
	m := map[int]int{
		1: 10,
		3: 13,
		4: 14,
	}
	_, ok := m[1]
	if ok {
		fmt.Println(m[1]) // 10
	}

	if value, inMap := m[2]; inMap {
		fmt.Println(value) // Условие не выполняется
	}

	// итерация по мапу
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(len(m))

}
