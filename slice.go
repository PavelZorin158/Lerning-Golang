package main

import "fmt"

func main() {
	// удаление 3 элемента среза
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl)
	// развертывание среза ...
	//т.к. 2 и последубщими аргументами должны быть элементы, а не срез
	sl = append(sl[:2], sl[3:]...)
	fmt.Println(sl)

	for i, s := range sl {
		fmt.Println(i, " - ", s)
	}
}
