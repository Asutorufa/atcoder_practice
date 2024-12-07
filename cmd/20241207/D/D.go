package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(countNumbers(N))
}

func countNumbers(n int) int {
	c := 0
	limit := math.Sqrt(float64(n))

	prime := make([]int, int(limit)+1)

	for i := 1; i <= int(limit); i++ {
		prime[i] = i
	}

	for i := 2; i <= int(limit); i++ {
		if prime[i] == i {
			for j := i * i; j <= int(limit); j += i {
				if prime[j] == j {
					prime[j] = i
				}
			}
		}
	}

	for i := 2; i <= int(limit); i++ {
		p := prime[i]
		q := prime[i/prime[i]]

		if p*q == i && q != 1 && p != q {
			c++
		} else if prime[i] == i {
			if math.Pow(float64(i), 8) <= float64(n) {
				c++
			}
		}
	}

	return c
}
