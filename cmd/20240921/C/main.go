package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unsafe"
)

var ABC = []byte("ABC")
var ans int

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, count int
	var ss string
	fmt.Fscan(r, &n, &count, &ss)

	s := unsafe.Slice(unsafe.StringData(ss), n)
	ans = bytes.Count(s, ABC)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for index := 0; index < count; index++ {
		var i int
		var byte string
		fmt.Fscan(r, &i, &byte)

		i--

		if s[i] != byte[0] {
			beforeStr(s, n, i)
			s[i] = byte[0]
			afterStr(s, n, i)
		}

		fmt.Fprintf(w, "%d\n", ans)
	}
}

func beforeStr(s []byte, length, i int) {
	if i >= 2 {
		if bytes.Equal(s[i-2:i+1], ABC) {
			ans--
			return
		}
	}

	if i > 0 && i+1 < length {
		if bytes.Equal(s[i-1:i+2], ABC) {
			ans--
			return
		}
	}

	if i+2 < length {
		if bytes.Equal(s[i:i+3], ABC) {
			ans--
			return
		}
	}
}

func afterStr(ns []byte, length, i int) {
	if i >= 2 {
		if bytes.Equal(ns[i-2:i+1], ABC) {
			ans++
			return
		}
	}

	if i > 0 && i+1 < length {
		if bytes.Equal(ns[i-1:i+2], ABC) {
			ans++
			return
		}
	}

	if i+2 < length {
		if bytes.Equal(ns[i:i+3], ABC) {
			ans++
			return
		}
	}
}
