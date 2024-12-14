package main

import "fmt"

func main() {
	var N, R int
	fmt.Scanln(&N, &R)

	for i := 0; i < N; i++ {
		var D, A int
		fmt.Scanln(&D, &A)

		switch D {
		case 1:
			if R >= 1600 && R <= 2799 {
				R = R + A
			}
		case 2:
			if R >= 1200 && R <= 2399 {
				R = R + A
			}
		}
	}

	fmt.Println(R)
}
