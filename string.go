package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "это строка"
	fmt.Printf("второе слово в кавычках \"%v\"\n", s[7:])
	s += " и еще"
	fmt.Println(s)
	bs := []byte("Это байтовый срез")
	fmt.Printf("%v\n", bs)
	fmt.Printf("%s\n", bs)

	//считывает строку с пробелами до \n
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// удаляем возврат каретки и перевод строки
	text = strings.Trim(text, "\r\n")
	fmt.Println("просто текст:\t", text, "\t", len(text), " элементов")
	// преобразовывает в срез рун
	textRune := []rune(text)
	fmt.Println("срез rune\t", string(textRune), "\t", len(textRune), " элементов")
	// преобразовываем в срез символов string
	textStr := strings.Split(text, "")
	fmt.Println("срез символов\t", textStr, "\t", len(textStr), "элементов")
}
