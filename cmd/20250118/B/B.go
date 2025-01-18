package main

import "fmt"

func main() {
	var X int
	fmt.Scanln(&X)

	sum := 1
	for i := 1; true; i++ {
		sum *= i

		if sum == X {
			fmt.Println(i)
			break
		}
	}
}
