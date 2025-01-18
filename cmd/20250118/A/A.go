package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	a, _ := strconv.Atoi(S[0:1])
	b, _ := strconv.Atoi(S[2:3])

	fmt.Println(a * b)
}
