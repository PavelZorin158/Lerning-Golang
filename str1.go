/*
На вход подается строка. Нужно определить, является ли она
правильной или нет. Правильная строка начинается с заглавной
буквы и заканчивается точкой. Если строка правильная -
вывести Right иначе - вывести Wrong
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//text := "Быть или не быть."
	textRune := []rune(text)
	if unicode.IsUpper(textRune[0]) && string(text[len(text)-1]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	// а еще можно выводить так
	os.Stdout.WriteString(string(textRune))
}
