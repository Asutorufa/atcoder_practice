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

	for i := len(line) - 1; i >= 0 && D > 0; i-- {
		if line[i] == '@' && D > 0 {
			line[i] = '.'
			D--
		}
	}

	fmt.Println(string(line))
}
