package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var remain atomic.Int32

func main1() {
	r := bufio.NewScanner(os.Stdin)

	r.Scan()
	wallsInit := strings.Fields(r.Text())

	h, err := strconv.Atoi(wallsInit[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(wallsInit[1])
	if err != nil {
		panic(err)
	}
	q, err := strconv.Atoi(wallsInit[2])
	if err != nil {
		panic(err)
	}

	remain.Store(int32(w * h))

	breakWall := make([][2]int, q)

	for i := range breakWall {
		r.Scan()
		vv := strings.Fields(r.Text())
		breakWall[i][0], err = strconv.Atoi(vv[0])
		if err != nil {
			panic(err)
		}
		breakWall[i][1], err = strconv.Atoi(vv[1])
		if err != nil {
			panic(err)
		}

		breakWall[i][0]--
		breakWall[i][1]--
	}

	walls := make([][]int, h)
	for i := range walls {
		walls[i] = make([]int, w)
	}

	wg := sync.WaitGroup{}
	for _, v := range breakWall {
		x, y := v[0], v[1]
		if walls[x][y] == 0 {
			walls[x][y] = 1

			remain.Add(-1)
			continue
		}

		wg.Add(4)

		go func() {
			defer wg.Done()
			breakWallFuncUP(walls, x, y)
		}()

		go func() {
			defer wg.Done()
			breakWallFuncDOWN(walls, x, y, h)
		}()

		go func() {
			defer wg.Done()
			breakWallFuncLEFT(walls, x, y)
		}()

		go func() {
			defer wg.Done()
			breakWallFuncRIGHT(walls, x, y, w)
		}()

		wg.Wait()
	}

	fmt.Println(remain.Load())
}

func breakWallFuncUP(walls [][]int, nowx, nowy int) {
	for {
		if nowx > 0 {
			nowx--
			if walls[nowx][nowy] == 0 {
				walls[nowx][nowy] = 1
				remain.Add(-1)
				break
			}
		} else {
			break
		}
	}
}

func breakWallFuncDOWN(walls [][]int, nowx, nowy, h int) {
	for {
		if nowx < h-1 {
			nowx++
			if walls[nowx][nowy] == 0 {
				walls[nowx][nowy] = 1
				remain.Add(-1)
				break
			}
		} else {
			break
		}
	}
}

func breakWallFuncLEFT(walls [][]int, nowx, nowy int) {
	for {
		if nowy > 0 {
			nowy--
			if walls[nowx][nowy] == 0 {
				walls[nowx][nowy] = 1
				remain.Add(-1)
				break
			}
		} else {
			break
		}
	}
}

func breakWallFuncRIGHT(walls [][]int, nowx, nowy, w int) {
	for {
		if nowy < w-1 {
			nowy++
			if walls[nowx][nowy] == 0 {
				walls[nowx][nowy] = 1
				remain.Add(-1)
				break
			}
		} else {
			break
		}
	}
}
