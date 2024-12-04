package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	s := store{max: N - 1}

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		s.AddOne(x-1, y-1)
	}

	fmt.Println(N*N - s.Size())
}

type store struct {
	max int
	x   map[int]map[int]struct{}
}

func (s *store) add(i, j int) {
	if s.x == nil {
		s.x = make(map[int]map[int]struct{})
	}

	if s.x[i] == nil {
		s.x[i] = make(map[int]struct{})
	}

	s.x[i][j] = struct{}{}
}

func (s *store) AddOne(i, j int) {
	s.add(i, j)

	for _, v := range ICan {
		if overflow(s.max, i+v.x, j+v.y) {
			continue
		}

		s.add(i+v.x, j+v.y)
	}
}

func (s *store) Size() int {
	cc := 0

	for i := range s.x {
		cc += len(s.x[i])
	}

	return cc
}

type XY struct {
	x, y int
}

var ICan = []XY{
	{2, 1},
	{1, 2},
	{-1, 2},
	{-2, 1},
	{-2, -1},
	{-1, -2},
	{1, -2},
	{2, -1},
}

func overflow(max int, i, j int) bool {
	if i < 0 || j < 0 || i > max || j > max {
		return true
	}

	return false
}
