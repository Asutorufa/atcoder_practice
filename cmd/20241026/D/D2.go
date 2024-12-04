package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	d := make([]int, M)
	for i := range d {
		d[i] = 1
	}

	for i := 0; i < N; i++ {
		var L, R int
		fmt.Fscan(br, &L, &R)

		d[R-1] = max(d[R-1], L+1)
	}

	// fmt.Println(d)

	for i := range d {
		if i == 0 {
			continue
		}
		d[i] = max(d[i], d[i-1])
	}

	ans := 0

	for i := range d {
		ans += (i + 1) - d[i] + 1
	}

	// fmt.Println(d, ans)

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
