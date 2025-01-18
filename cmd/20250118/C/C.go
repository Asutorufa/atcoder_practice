package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscanln(br, &Q)

	array := make([]int, 0, 1000000)

	scan := bufio.NewScanner(br)

	sub := 0
	for i := 0; i < Q; i++ {
		scan.Scan()

		fields := strings.Fields(scan.Text())

		switch fields[0] {
		case "1":
			lenght, _ := strconv.Atoi(fields[1])
			last := 0
			if len(array) > 0 {
				last = array[len(array)-1]
			}
			array = append(array, last+lenght)

		case "2":
			sub = array[0]
			array = array[1:]
			if len(array) == 0 {
				sub = 0
			}

		case "3":
			index, _ := strconv.Atoi(fields[1])

			if index == 1 {
				fmt.Println(0)
			} else {
				fmt.Println(array[index-2] - sub)
			}
		}
	}
}
