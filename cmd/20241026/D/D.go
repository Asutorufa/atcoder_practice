package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	store := store{}

	for i := 0; i < N; i++ {
		var L, R int
		fmt.Fscan(br, &L, &R)

		for _, r := range fixed(M, L, R) {
			store.add(r.L, r.R)
		}
	}

	/*
		1,1 1,2 1,3 1,4 3
		2,1 2,2 2,3 2,4 1
		3,1 3,2 3,3 3,4 1
		4,1 4,2 4,3 4,4 4
	*/
	// fmt.Println(store.Size(), ((M*M)/2+M/2+M%2)-store.Size())
	// fmt.Println(store.x)
	fmt.Println(((M*M)/2 + M/2 + M%2) - store.Size())
}

type XY struct {
	L, R int
}

func fixed(max int, l, r int) []XY {
	resp := []XY{}
	resp = append(resp, XY{l, r})

	for i := max - r; i >= 0; i-- {
		for j := l; j > 0; j-- {
			if j > r+i {
				continue
			}

			resp = append(resp, XY{j, r + i})
		}
	}

	return resp
}

type store struct {
	x map[int]map[int]struct{}
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

func (s *store) Size() int {
	cc := 0

	for i := range s.x {
		cc += len(s.x[i])
	}

	return cc
}
