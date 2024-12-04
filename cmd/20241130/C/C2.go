package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	br := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscanln(br, &N, &M)

	A := readInt(N, br)

	found := func(x int) int {
		for i := range A {
			if x >= A[i] {
				return i + 1
			}
		}

		return -1
	}

	for _, v := range readInt(M, br) {
		fmt.Println(found(v))
	}
}

func readInt(n int, r *bufio.Reader) []int {
	var resp = make([]int, 0, n)
	readIntFunc(n, r, func(x, _ int) { resp = append(resp, x) })
	return resp
}
