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

	lastTime := 0
	current := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Fscan(br, &A, &B)

		if lastTime > 0 {
			moreru := A - lastTime
			if current < moreru {
				current = 0
			} else {
				current -= moreru
			}
		}
		lastTime = A
		current += B
	}

	fmt.Println(current)
}

func readInt(n int, r *bufio.Reader) []int {
	var resp = make([]int, 0, n)
	readIntFunc(n, r, func(x, _ int) { resp = append(resp, x) })
	return resp
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}
