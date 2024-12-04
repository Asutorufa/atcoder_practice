package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscanln(br, &N, &M)

	var R = 200010
	ids := make([]int, R)
	for i := 0; i < R; i++ {
		ids[i] = -1
	}

	readIntFunc(N, br, func(x int, index int) {
		for R > x {
			R--
			ids[R] = index + 1
		}
	})

	readIntFunc(M, br, func(x, _ int) {
		fmt.Println(ids[x])
	})
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}

/*
1 1
3
5

*/
