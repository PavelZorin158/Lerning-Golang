// отражает слово

package main

import (
	"fmt"
	"strings"
)

func main() {
	var in, out string
	fmt.Scan(&in)
	sliceIn := strings.Split(in, "")
	n := len(sliceIn)
	for i := 0; i < n/2; i++ {
		sliceIn[i], sliceIn[n-1-i] = sliceIn[n-1-i], sliceIn[i]
	}
	out = strings.Join(sliceIn, "")
	fmt.Println(out)
}
