package main

import (
	"bufio"
	"fmt"
	"os"
)

// e，直接暴力解就可以，三循环，想复杂了
// 因为有两个加湿器，同时计算每个床放置加湿器后最大加湿范围，暴力解，感觉实用性不高
func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W, D int
	fmt.Fscanln(br, &H, &W, &D)

	scan := bufio.NewScanner(br)

	mas := make([]string, H)
	for i := 0; i < H; i++ {
		scan.Scan()
		mas[i] = scan.Text()
	}

	ans := 0
	for i1, s := range mas {
		for j1, v := range s {
			if v == '#' {
				continue
			}

			for i2, s2 := range mas {
				for j2, v2 := range s2 {
					if v2 == '#' || (i1 == i2 && j1 == j2) {
						continue
					}

					tmp := 0

					for i, s3 := range mas {
						for j, v3 := range s3 {
							if v3 == '.' && (Abs(i-i1)+Abs(j-j1) <= D || Abs(i-i2)+Abs(j-j2) <= D) {
								tmp++
							}
						}
					}

					ans = max(ans, tmp)
				}
			}
		}
	}

	fmt.Println(ans)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
