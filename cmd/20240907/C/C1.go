package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main1() {
	r := bufio.NewScanner(os.Stdin)

	r.Scan()
	s := r.Bytes()
	r.Scan()
	t := r.Bytes()

	n := len(s)

	ans := []string{}
	for !bytes.Equal(s, t) {
		nxt := make([]byte, n)
		for i := range nxt {
			nxt[i] = 'z'
		}

		for i := 0; i < n; i++ {
			if s[i] != t[i] {
				tmp := clone(s)
				tmp[i] = t[i]
				nxt = min(nxt, tmp)
			}
		}

		fmt.Println("-", string(nxt))
		ans = append(ans, string(nxt))
		s = nxt
	}

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Println(v)
	}
}

func min(x, y []byte) []byte {
	if string(x) <= string(y) {
		return x
	}
	return y
}

func clone(x []byte) []byte {
	n := make([]byte, len(x))
	copy(n, x)
	return n
}
