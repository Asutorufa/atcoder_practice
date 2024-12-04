package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	x := map[byte]struct{}{
		'A': {},
		'B': {},
		'C': {},
	}

	l, _, _ := br.ReadLine()

	for _, v := range l {
		delete(x, v)
	}

	if len(x) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
