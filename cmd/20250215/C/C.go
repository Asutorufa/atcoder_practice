package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	br := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscanln(br, &N, &M)

	var cache = map[Point]bool{}
	var count int
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscanln(br, &x, &y)

		if x == y {
			count++
			continue
		}

		if cache[Point{x, y}] || cache[Point{y, x}] {
			count++
			continue
		}

		cache[Point{x, y}] = true
	}

	fmt.Println(count)
}
