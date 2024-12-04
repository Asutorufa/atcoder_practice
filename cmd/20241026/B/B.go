package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	x := [8][8]bool{}
	for i := 0; i < 8; i++ {
		l, _, _ := br.ReadLine()
		for j, v := range l {
			if v == '#' {
				for z := range x[i] {
					x[i][z] = true
					x[z][j] = true
				}
			}
		}
	}

	cc := 0
	for i := range x {
		for j := range x[i] {
			if !x[i][j] {
				cc++
			}
		}
	}

	fmt.Println(cc)
}
