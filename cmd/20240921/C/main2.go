package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main2() {
	r := bufio.NewReader(os.Stdin)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, count int
	var s string
	fmt.Fscan(r, &n, &count, &s)

	str := []byte(s)

	for index := 0; index < count; index++ {
		var i int
		var ss string
		fmt.Fscan(r, &i, &ss)

		str[i-1] = ss[0]

		fmt.Fprintf(w, "%d\n", bytes.Count(str, []byte("ABC")))
	}
}
