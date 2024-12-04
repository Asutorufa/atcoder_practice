package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	lines, _, _ := br.ReadLine()

	store := map[byte]int{}
	for _, v := range lines {
		store[v] = store[v] + 1
	}

	if store['1'] == 1 && store['2'] == 2 && store['3'] == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
