package main

import (
	"fmt"
	"strings"
)

func main() {
	var s int

	fmt.Scanf("%d", &s)

	var a []int

	for i := 0; i <= 10; i++ {
		for j := 0; j < s%3; j++ {
			a = append(a, i)
		}

		s /= 3
	}

	fmt.Println(len(a))
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}
