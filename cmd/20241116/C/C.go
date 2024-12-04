package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscanln(br, &N, &K)

	var S string
	fmt.Fscanln(br, &S)

	KSub1Index := 0

	kStart, kEnd := 0, 0

	currentBlock := 0

	lastByte := '0'
	for index, v := range S {
		if v == '1' && lastByte == '0' {
			currentBlock++

			if currentBlock == K {
				kStart = index
			}
		}

		if v == '0' && lastByte == '1' {
			if currentBlock == K-1 {
				KSub1Index = index
			}

			if currentBlock == K {
				kEnd = index
				lastByte = v
				break
			}
		}

		lastByte = v
	}

	if lastByte == '1' {
		kEnd = len(S)
	}

	str := string(S[:KSub1Index])
	str += string(S[kStart:kEnd])
	str += strings.Repeat("0", kStart-KSub1Index)
	str += string(S[kEnd:])

	fmt.Println(str)
}

/*
15 3
010011100111001
*/
