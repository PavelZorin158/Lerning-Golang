package main

import "fmt"

func main() {
	var yy int
	var x, y, p float64
	fmt.Scan(&x, &p, &y)
	for yy = 0; x < y; yy++ {
		x += x / 100 * p // сумма годовых процентов
		x = float64(int(x*100) / 100)
	}
	fmt.Println(yy)
}
