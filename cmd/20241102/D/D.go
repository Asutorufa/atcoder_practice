package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var currentCache = map[point]bool{}

var mas [][]int
var H, W, K int
var count int

func main() {
	br := bufio.NewReader(os.Stdin)

	fmt.Fscanln(br, &H, &W, &K)

	mas = make([][]int, H)
	for i := 0; i < H; i++ {
		mas[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		lines, _, _ := br.ReadLine()
		for j, v := range lines {
			if v == '#' {
				mas[i][j] = -1
			}
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// fmt.Println(i, j, cache)
			if mas[i][j] == 0 {
				currentCache = map[point]bool{}
				move(i, j, 0)
			}
		}
	}

	fmt.Println(count)
}

func move(x, y, step int) bool {
	if currentCache[point{x, y}] {
		return false
	}

	if mas[x][y] == -1 {
		return false
	}

	if step == K {
		count++
		return false
	}

	moveUP(x, y, step)
	moveDOWN(x, y, step)
	moveLEFT(x, y, step)
	moveRIGHT(x, y, step)

	return false
}

func moveUP(x, y, step int) bool {
	if x == 0 {
		return false
	}

	currentCache[point{x, y}] = true
	move(x-1, y, step+1)
	currentCache[point{x, y}] = false

	return true
}

func moveDOWN(x, y, step int) bool {
	if x == H-1 {
		return false
	}

	currentCache[point{x, y}] = true
	move(x+1, y, step+1)
	currentCache[point{x, y}] = false
	return true
}

func moveLEFT(x, y, step int) bool {
	if y == 0 {
		return false
	}

	currentCache[point{x, y}] = true
	move(x, y-1, step+1)
	currentCache[point{x, y}] = false
	return true
}

func moveRIGHT(x, y, step int) bool {
	if y == W-1 {
		return false
	}

	currentCache[point{x, y}] = true
	move(x, y+1, step+1)
	currentCache[point{x, y}] = false
	return true
}
