package main

import "fmt"

//обьявление структуры
type person struct {
	name string
	age  int
}

// встраивоемый тип
type android struct {
	Person person
	Model  string
}

// встраивоемый тип с анонимным полем
type android2 struct {
	person
	Model string
}

// метод типа person
func (c *person) lenName() int {
	return len(c.name)
}

func main() {

	rick := person{"Rick", 47}

	Kat := android{
		Model: "R2D2",
		Person: person{
			name: "Katrin",
			age:  46,
		},
	}

	tima := android2{
		person: person{
			name: "Tima",
			age:  12},
		Model: "kinder"}
	fmt.Print(rick.name, rick.age)
	fmt.Println(" длинна имени ", rick.lenName(), " букв")
	fmt.Println(Kat.Person.name, " ", Kat.Person.lenName())
	// в типе с анонимным полем доступны сразу поля анонимного типа
	fmt.Println(tima.person.name, "'или так'", tima.name)
}
