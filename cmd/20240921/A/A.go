package main

import (
	"bufio"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		a, err := r.ReadByte()
		if err != nil {
			break
		}

		if a == '.' {
			continue
		}

		if a == '\n' {
			break
		}

		os.Stdout.Write([]byte{a})
	}

	os.Stdout.Write([]byte("\n"))
}
