package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	r.Scan()

	hss := strings.Fields(r.Text())

	hs := make([]int, len(hss))

	var err error
	for i := range hss {
		hs[i], err = strconv.Atoi(hss[i])
		if err != nil {
			panic(err)
		}
	}

	counts := []string{}
	for i := range hs {
		count := 0
		for j := i + 1; j < len(hs); j++ {
			if hasNoBigger(hs[i+1:j+1], hs[j]) {
				// fmt.Println(j, hs[i], hs[j])
				count++
			}
		}

		counts = append(counts, fmt.Sprint(count))
	}

	fmt.Println(strings.Join(counts, " "))
}

func hasNoBigger(hs []int, i int) bool {
	// fmt.Println(hs, i)
	for z := range hs {
		if hs[z] > i {
			return false
		}
	}

	return true
}
