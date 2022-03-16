package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Hour(), ":", now.Minute())

	// преобразование time в string
	fmt.Println(now.Format("(Monday) 2-January-2006 15:04:05"))

	//преобразование string в time
	firstTime, err := time.Parse("2006/01/02 15:04", "2022/03/16 17:45")
	if err != nil {
		panic(err)
	}

	// LoadLocation находит временную зону в справочнике IANA
	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Europe/Minsk")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)

	// парсит дату и время в строковом представлении с отдельным указанием временной зоны
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 22 05:45:10pm", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format("02-01-2006 15:04:05"))
	fmt.Println(secondTime.Format("02-01-2006 15:04:05"))

	// true если firstTime > secondTime
	fmt.Println(firstTime.Before(secondTime))
	// true если равны
	fmt.Println(firstTime.Equal(secondTime))

	future := now.Add(time.Hour * 5) // перемещаемся на 12 часов вперед
	past := now.AddDate(-1, -2, -3)
	// перемещаемся на 1 год, два месяца и 3 дня назад
	fmt.Println(future.Format("02-01-2006 15:04:05"))
	fmt.Println(past.Format("02-01-2006 15:04:05"))

	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past))
}
