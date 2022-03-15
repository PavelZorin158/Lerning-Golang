package main

import "fmt"

type bat struct {
	power int
}

// или так, вариант.2
// type bat int

func (b bat) String() string {
	var s = "["
	for i := 10; i > 0; i-- {
		if b.power >= i {
			// if int(b) >= i {
			s += "X"
		} else {
			s += " "
		}
	}
	s += "]"
	return s
}

func main() {
	var s string
	var p int //заряд батареи
	var batteryForTest bat
	fmt.Scan(&s)
	for _, v := range s {
		if v == '1' {
			p++
		}
	}
	batteryForTest.power = p
	// batteryForTest = bat(p)
	fmt.Println(batteryForTest)
}
