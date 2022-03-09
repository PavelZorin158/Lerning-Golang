//Дано натуральное число A > 1. Определите, каким по счету
//числом Фибоначчи оно является, то есть выведите такое число n,
//что φn=A. Если А не является числом Фибоначчи, выведите число -1.

package main

import "fmt"

func main() {
	// 0 1 1 2 3 5 8
	var in, count, f1, f2, f3 int
	f2 = 1
	f3 = f1 + f2
	fmt.Scan(&in)
	for count = 2; f3 < in; count++ {
		f1, f2, f3 = f2, f3, f2+f3
	}
	if f3 == in {
		fmt.Println(count)
	} else {
		fmt.Println(-1)
	}
}
