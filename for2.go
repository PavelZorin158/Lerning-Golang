package main

// ПЕРЕБОР ЦИФР ЧИСЛА
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)

	// 2 ВАРИАНТ
	var A, B, C int
	fmt.Scanf("%1d%1d%1d", &A, &B, &C)
	fmt.Println(A + B + C)
}
