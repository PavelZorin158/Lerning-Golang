package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T (byte)\n", data)
	// string умеет собирать из []byte русские буквы
	fmt.Println(string(data))
	fmt.Println("----------------")

	r := []rune(string(data))
	fmt.Printf("%T (rune)\n", r)
	fmt.Println(string(r))
	fmt.Println("----------------")

	s := string(data)
	slic := strings.Split(s, "\r\n") // разбиваем по строкам
	fmt.Println(strings.Join(slic, "/"))

	s = "Вот твкой текст.\r\nтра-ля-ля\r\nTelepuziki"
	data = []byte(s)
	// permission разрешения. смотреть png в папке
	err = os.WriteFile("test_out.txt", data, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

}
