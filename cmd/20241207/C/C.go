package main

import (
	"bufio"
	"fmt"
	"os"
)

var H, W, D int

func main() {
	br := bufio.NewReaderSize(os.Stdin, 1024*1024)

	fmt.Fscanln(br, &H, &W, &D)

	scan := bufio.NewScanner(br)
	buf := make([]byte, 1024*1024)
	scan.Buffer(buf, bufio.MaxScanTokenSize)
	scan.Split(bufio.ScanWords)

	masu := make([]string, H)
	for i := 0; i < H; i++ {
		scan.Scan()
		masu[i] = scan.Text()
	}

	for x, v := range masu {
		for y, vv := range v {
			if vv == 'H' {
				queue = append(queue, [3]int{x, y, D})
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		run2(masu, p[0], p[1], p[2])
	}

	fmt.Println(len(used))
}

var queue = [][3]int{}
var used = map[[2]int]int{}

// run 递归同样可解，但递归函数调用太费时间。。。无语
func run(masu []string, x, y int, left int) {
	if masu[x][y] == '#' {
		return
	}

	if left != D && masu[x][y] == 'H' {
		return
	}

	p := [2]int{x, y}

	if fu, ok := used[p]; ok && fu >= left {
		return
	}

	used[p] = left

	if left == 0 {
		return
	}

	if x > 0 {
		run(masu, x-1, y, left-1)
	}

	if y > 0 {
		run(masu, x, y-1, left-1)
	}

	if x < len(masu)-1 {
		run(masu, x+1, y, left-1)
	}

	if y < len(masu[0])-1 {
		run(masu, x, y+1, left-1)
	}
}

func run2(masu []string, x, y int, left int) {
	if masu[x][y] == '#' {
		return
	}

	if left != D && masu[x][y] == 'H' {
		return
	}

	p := [2]int{x, y}

	if fu, ok := used[p]; ok && fu >= left {
		return
	}

	used[p] = left

	if left == 0 {
		return
	}

	if x > 0 {
		queue = append(queue, [3]int{x - 1, y, left - 1})
	}

	if y > 0 {
		queue = append(queue, [3]int{x, y - 1, left - 1})
	}

	if x < len(masu)-1 {
		queue = append(queue, [3]int{x + 1, y, left - 1})
	}

	if y < len(masu[0])-1 {
		queue = append(queue, [3]int{x, y + 1, left - 1})
	}
}
