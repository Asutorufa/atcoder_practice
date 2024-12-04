package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scanln(&A, &B, &C, &D)

	x := []int{A, B, C, D}

	z := map[int]int{}

	count := 0
	for _, v := range x {
		if z[v] == 1 {
			count++
			delete(z, v)
			continue
		}

		z[v] = z[v] + 1
	}

	fmt.Println(count)
}
