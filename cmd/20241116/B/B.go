package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	lines, _, _ := br.ReadLine()

	fields := strings.FieldsFunc(string(lines[1:len(lines)-1]), func(r rune) bool {
		return r == '|'
	})

	for _, v := range fields {
		fmt.Print(len(v), " ")
	}
	fmt.Println()
}
