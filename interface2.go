package main

import "fmt"

type car struct {
	model string
	yar   int
}

// добавляем к типу метод:  string() string , который подходит
//под интерфейс fmt.Println() и поэтому он срабатывает в fmt.Println
func (c car) String() string {
	return fmt.Sprintf("%v (%d)", c.model, c.yar)
}

func main() {
	myCar := car{model: "Lamba", yar: 2020}
	fmt.Println("моя тачка - ", myCar)
}
