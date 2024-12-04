package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewScanner(os.Stdin)

	r.Scan()
	s := bytes.TrimSpace(r.Bytes())
	r.Scan()
	t := bytes.TrimSpace(r.Bytes())

	x := []int{}

	for i := range s {
		if s[i] > t[i] {
			x = append(x, i)
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < t[i] {
			x = append(x, i)
		}
	}

	fmt.Println(len(x))

	for _, v := range x {
		s[v] = t[v]
		fmt.Println(string(s))
	}

}
