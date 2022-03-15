// приведение типа в пустом интерфейсе

package main

import "fmt"

func main() {
	//value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	var value1, value2, operation interface{} = float64(12), float64(10), "#"
	var v1, v2, res float64
	var op string
	var ok bool
	if v1, ok = value1.(float64); !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	if v2, ok = value2.(float64); !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}
	if op, ok = operation.(string); !ok {
		fmt.Println("неизвестная операция")
		return
	}
	switch op {
	case "+":
		res = v1 + v2
	case "-":
		res = v1 - v2
	case "/":
		res = v1 / v2
	case "*":
		res = v1 * v2
	default:
		fmt.Println("неизвестная операция")
		return
	}
	fmt.Printf("%.4f", res)
}
