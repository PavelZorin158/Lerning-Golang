package main

import "fmt"

func main() {
	// удаление 3 элемента среза
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl)
	// развертывание среза ...
	sl = append(sl[:2], sl[3:]...)
	fmt.Println(sl)
}
