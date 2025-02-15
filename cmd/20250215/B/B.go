package main

import "fmt"

func main() {
	var S string
	fmt.Scanln(&S)

	var count int
	for i := 0; i < 100; i++ {
		for j := 0; j < len(S); j++ {
			ii := j
			jj := ii + i
			kk := jj + i

			if jj >= len(S) || kk >= len(S) {
				break
			}

			if S[ii] == 'A' && S[jj] == 'B' && S[kk] == 'C' {
				count++
			}
		}
	}

	fmt.Println(count)
}
