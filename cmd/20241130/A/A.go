package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, D int
	fmt.Fscanln(br, &N, &D)

	line, _, _ := br.ReadLine()

	count := 0

	for _, v := range line {
		if v == '.' {
			count++
		}
	}

	fmt.Println(count + D)
}
