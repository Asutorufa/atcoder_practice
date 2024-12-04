package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	remain := w * h

	walls := newWalls(h, w)

	for ii := 1; ii <= q; ii++ {
		r.Scan()
		vv := strings.Fields(r.Text())
		x, err := strconv.Atoi(vv[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(vv[1])
		if err != nil {
			panic(err)
		}
		x--
		y--

		// fmt.Println("---")
		// walls.Print()
		// fmt.Println("---")

		// v1 := walls.y[y].LowerBound(x)
		// v2 := walls.x[x].LowerBound(y)
		// var v1v, v2v int

		// if v1 != -1 {
		// 	v1v = walls.y[y].get(v1)
		// }

		// if v2 != -1 {
		// 	v2v = walls.x[x].get(v2)
		// }

		// if v1v == x && v2v == y {
		// 	walls.x[x].removeByIndex(v2)
		// 	walls.y[y].removeByIndex(v1)
		// 	remain--
		// } else {
		// 	v3 := walls.y[y].ReverseLowerBound(x - 1)
		// 	v4 := walls.x[x].ReverseLowerBound(y - 1)

		// 	fmt.Println(v1, v2, v3, v4)

		// 	if v1 != -1 {
		// 		walls.y[y].removeByIndex(v1)
		// 		walls.x[v1v].removeByValue(y)
		// 		remain--
		// 	}

		// 	if v3 != -1 {
		// 		v := walls.y[y].get(v3)
		// 		walls.y[y].removeByIndex(v3)
		// 		walls.x[v].removeByValue(y)
		// 		remain--
		// 	}

		// 	if v2 != -1 {
		// 		walls.x[x].removeByIndex(v2)
		// 		walls.y[v2v].removeByValue(x)
		// 		remain--
		// 	}

		// 	if v4 != -1 {
		// 		v := walls.x[x].get(v4)
		// 		walls.x[x].removeByIndex(v4)
		// 		walls.y[v].removeByValue(x)
		// 		remain--
		// 	}

		// }

		if walls.remove(x, y) {
			remain--
			continue
		}

		if walls.removeUP(x, y) {
			remain--
		}

		if walls.removeDOWN(x, y) {
			remain--
		}

		if walls.removeLEFT(x, y) {
			remain--
		}

		if walls.removeRIGHT(x, y) {
			remain--
		}
	}

	// fmt.Println("---")
	// walls.Print()
	fmt.Println(remain)
}

type set struct {
	slice []int
	len   int
}

func (s *set) add(v int) {
	s.slice = append(s.slice, v)
	s.len++
}

func (s *set) removeByIndex(i int) {
	if i < 0 || i >= s.len {
		return
	}

	copy(s.slice[i:], s.slice[i+1:])
	s.len--
	s.slice = s.slice[:s.len]
}

func (s *set) removeByValue(i int) {
	x := s.findByValue(i)
	if x == -1 {
		// fmt.Println("removeByValue error", i, *s)
		return
	}

	// fmt.Println("removeByValue", x)
	s.removeByIndex(x)
}

func (s *set) findByValue(x int) int {
	i, found := sort.Find(s.len, func(v int) int {
		if s.slice[v] > x {
			return -1
		} else if s.slice[v] < x {
			return 1
		}

		return 0
	})
	if found {
		return i
	}
	return -1
}

func (s *set) LowerBound(v int) int {
	i := sort.Search(s.len, func(i int) bool { return s.slice[i] >= v })

	if i == s.len {
		return -1
	}

	return i
}

func (s *set) ReverseLowerBound(v int) int {
	max := -1
	sort.Search(s.len, func(i int) bool {
		if s.slice[i] <= v {
			max = i
		}
		return s.slice[i] > v
	})

	return max
}

func (s *set) prev(i int) int {
	if i == -1 && s.len != 0 {
		return s.len - 1
	}

	if i <= 0 || i >= s.len {
		return -1
	}

	return i - 1
}

func (s *set) get(i int) int {
	return s.slice[i]
}

type walls struct {
	x []set
	y []set
}

func newWalls(h, w int) walls {
	wls := walls{
		x: make([]set, h),
		y: make([]set, w),
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			wls.x[j].add(i)
			wls.y[i].add(j)
		}
	}

	return wls
}

func (w *walls) remove(x, y int) bool {
	i := w.x[x].findByValue(y)
	if i != -1 {
		// fmt.Println("remove", x, y)
		w.x[x].removeByIndex(i)
		w.y[y].removeByValue(x)
		return true
	}

	// fmt.Println("remove error", w.x[x], w.y[y], y)

	return false
}

func (w *walls) removeUP(x, y int) bool {
	i := w.y[y].LowerBound(x)

	pre := w.y[y].prev(i)

	if pre == -1 {
		return false
	}

	prev := w.y[y].get(pre)

	// fmt.Println("remove UP", x, w.y[y][pre])

	w.y[y].removeByIndex(pre)
	w.x[prev].removeByValue(y)
	return true
}

func (w *walls) removeDOWN(x, y int) bool {
	i := w.y[y].LowerBound(x)
	if i == -1 {
		return false
	}

	v := w.y[y].get(i)
	// fmt.Println("remove DOWN", x, y, i, w.y[y][i])
	w.y[y].removeByIndex(i)
	w.x[v].removeByValue(y)
	return true
}

func (w *walls) removeLEFT(x, y int) bool {
	i := w.x[x].LowerBound(y)

	pre := w.x[x].prev(i)
	if pre == -1 {
		return false
	}

	prev := w.x[x].get(pre)
	// fmt.Println("remove LEFT", w.x[x][pre], y)

	w.x[x].removeByIndex(pre)
	w.y[prev].removeByValue(x)
	return true
}

func (w *walls) removeRIGHT(x, y int) bool {
	i := w.x[x].LowerBound(y)
	if i == -1 {
		return false
	}

	v := w.x[x].get(i)
	// fmt.Println("remove RIGHT", w.x[x][i], y)
	w.x[x].removeByIndex(i)
	w.y[v].removeByValue(x)
	return true
}

func (w *walls) Print() {
	for i := 0; i < len(w.x); i++ {
		for j := 0; j < len(w.y); j++ {
			if w.x[i].findByValue(j) != -1 {
				fmt.Print(1, " ")
			} else {
				fmt.Print(0, " ")
			}
		}
		fmt.Println()
	}
}

/*

5 4 5 0 0
4 3 2 4 0
3 2 1 2 3
0 3 2 0 0
0 4 0 0 0

*/
