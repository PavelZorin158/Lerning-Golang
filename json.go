/*В сведениях о каждом студенте содержится информация о полученных
им оценках (Rating). Требуется прочитать данные, и рассчитать
среднее количество оценок, полученное студентами группы.
Ответ на задачу требуется записать на стандартный вывод
в формате JSON в следующей форме:*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type result struct {
		Average float64
	}
	type stud struct {
		LastName   string
		FirstName  string
		MiddleName string
		Birthday   string
		Address    string
		Phone      string
		Rating     []int
	}

	type univer struct {
		ID       int
		Number   string
		Year     int
		Students []stud
	}
	var s univer
	var ratN int
	var res result
	data, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println("чето пошло не так \n", err)
		return
	}
	for i := 0; i < len(s.Students); i++ {
		ratN += len(s.Students[i].Rating)
	}

	aver := float64(ratN) / float64(len(s.Students))
	aver = float64(int(aver*10)) / 10
	res.Average = aver
	data, err = json.MarshalIndent(res, "", "    ")
	fmt.Println(string(data))
}
