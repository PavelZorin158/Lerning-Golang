package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Scan(&in)
	sliceIn := strings.Split(in, "")
	n := len(sliceIn)
	res := "Палиндром"
	for i := 0; i < int(n/2); i++ {
		if sliceIn[i] != sliceIn[n-1-i] {
			res = "Нет"
			break
		}
	}
	fmt.Println(res)
}
