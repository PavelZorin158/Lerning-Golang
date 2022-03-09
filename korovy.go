// проверка на ошибку ввода
// выбор правильного окончания
package main

import (
	"fmt"
)

func main() {
	var in, n, scan int
	var s string

	scan, err := fmt.Scan(&in)
	if err != nil {
		fmt.Println("Ошибка ввода", err)
		fmt.Printf("scan : %v\n err : %v\n ", scan, err)
		return
	}

	n = in
	if n > 20 {
		n = n % 10
	}
	switch {
	case n == 1:
		s = "korova"
	case 1 < n && n < 5:
		s = "korovy"
	default:
		s = "korov"
	}
	fmt.Println(in, s)
}
