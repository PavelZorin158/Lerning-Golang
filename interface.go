package main

import "fmt"

// обьявляем структуру
type famely struct {
	name string
	age  int
}

// добавляем в нее метод
func (f famely) print() bool {
	if f.age <= 0 || f.age > 100 {
		return false
	}
	fmt.Println(f.name, " : ", f.age, " лет")
	return true
}

// обьявляем интерфейс
type fPrinter interface {
	print() bool
}

// создаем функцию, которая принемает данные типа интерфейс fPrint
func printFam(ff fPrinter) {
	if ok := ff.print(); !ok {
		fmt.Println(" странный возраст")
	}
	return
}

func main() {
	//создаем экземпляр типа famely
	var zorin = famely{name: "Павел", age: 0}
	//передаем функции экземпляр подходящего типа под интерфейс
	printFam(zorin)
}
