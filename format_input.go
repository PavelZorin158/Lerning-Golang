package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, i int
	// считываются по одной цифре (d-децимал)
	// 12 - в c,d ничего не записывается
	// 123456 - 56 улетает в следующий ввод
	fmt.Scanf("%1d%1d%1d%1d", &a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// считываем лишние цифры, чтоб освободить ввод
	fmt.Scan(&i)

	fmt.Println(i)
}
