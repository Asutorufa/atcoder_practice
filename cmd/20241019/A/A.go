package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, C int

	fmt.Fscanf(br, "%d %d\n", &N, &C)

	var Ts []int

	for i := 0; i < N; i++ {
		var T int
		fmt.Fscanf(br, "%d", &T)
		Ts = append(Ts, T)
	}

	last := Ts[0]
	count := 1

	for _, v := range Ts[1:] {
		if v-last >= C {
			count++
			last = v
		}
	}

	fmt.Println(count)
}
