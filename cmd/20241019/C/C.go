package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int64
	fmt.Fscanln(br, &N)

	As := make([]int64, 0, N)
	Bs := make([]int64, 0, N-1)
	minA, minB := int64(-1), int64(-1)

	for i := int64(0); i < N; i++ {
		var A int64
		fmt.Fscan(br, &A)
		if minA == -1 || A < minA {
			minA = A
		}

		As = append(As, A)
	}

	if N == 1 {
		fmt.Println(As[0])
		return
	}

	for i := int64(0); i < N-1; i++ {
		var B int64
		fmt.Fscan(br, &B)

		if minB == -1 || B < minB {
			minB = B
		}

		Bs = append(Bs, B)
	}

	if minB < minA {
		fmt.Println(-1)
		return
	}

	sort.Slice(As, func(i, j int) bool { return As[i] > As[j] })
	sort.Slice(Bs, func(i, j int) bool { return Bs[i] > Bs[j] })

	var remain int64
	for i, B := range Bs {
		x := i
		if remain != 0 {
			x += 1
		}

		if As[x] <= B {
			// fmt.Println("A <= B continue", "A", A, "B", B)
			continue
		} else {
			if remain != 0 {
				fmt.Println(-1)
				return
			}

			remain = As[x]
		}
	}

	if remain > 0 {
		fmt.Println(remain)
	} else {
		fmt.Println(As[len(As)-1])
	}
}
