package main

import (
	"fmt"
	"testing"
)

func TestD(t *testing.T) {
	a := make([]int, 0, 20)

	a = append(a, 1, 2, 3)

	x := append(a, 4, 5, 6, 7)
	fmt.Println(x)

	y := append(a, 9)
	fmt.Println(y)

	fmt.Println(x)
}
