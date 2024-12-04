package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewReader(os.Stdin)

	x, err := scan.Peek(3)
	if err != nil {
		fmt.Println("Invalid")
	}

	if x[0] == '1' && x[2] == '0' {
		fmt.Println("Yes")
	} else if x[0] == '0' && x[2] == '1' {
		fmt.Println("No")
	} else {
		fmt.Println("Invalid")
	}
}
