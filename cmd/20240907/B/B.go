package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := bufio.NewScanner(os.Stdin)

	x.Scan()
	n, err := strconv.Atoi(x.Text())
	if err != nil {
		panic(err)
	}

	elems := make([][]int, n)
	for i := range elems {
		elems[i] = make([]int, n)
	}

	for i := range elems {
		x.Scan()
		vv := strings.Fields(x.Text())

		for j, vr := range vv {
			z, err := strconv.Atoi(vr)
			if err != nil {
				panic(err)
			}

			elems[i][j] = z
		}
	}

	last := 0
	for i := range elems {
		if i == 0 {
			last = elems[i][0]
			continue
		}

		if last-1 >= i {
			last = elems[last-1][i]
		} else {
			last = elems[i][last-1]
		}
	}

	fmt.Println(last)
}
