package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	type pair struct {
		Q, R int
	}

	var pairs []pair
	for i := 0; i < N; i++ {
		var Q, R int
		fmt.Fscan(br, &Q, &R)
		pairs = append(pairs, pair{Q, R})
	}

	var Q int
	fmt.Fscan(br, &Q)

	for i := 0; i < Q; i++ {
		var T, D int
		fmt.Fscan(br, &T, &D)
		T--

		p := pairs[T]

		c := D % p.Q

		if c == p.R {
			fmt.Println(D)
			continue
		}

		if c > p.R {
			fmt.Println(D + (p.Q - c) + p.R)
			continue
		}

		fmt.Println(D + (p.R - c))
	}
}
