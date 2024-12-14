package main

import "fmt"

func main() {
	var N int
	var C1, C2 string
	fmt.Scanln(&N, &C1, &C2)

	var S string
	fmt.Scanln(&S)

	for _, v := range S {
		if v == rune(C1[0]) {
			fmt.Print(C1)
		} else {
			fmt.Print(C2)
		}
	}
	fmt.Println()
}
