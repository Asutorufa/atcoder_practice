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

	zz := map[int]int{}
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		if x, ok := zz[A]; ok {
			fmt.Print(x)
		} else {
			fmt.Print(-1)
		}

		if i != N-1 {
			fmt.Print(" ")
		}

		zz[A] = i + 1
	}
	fmt.Println()
}
