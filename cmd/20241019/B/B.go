package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var N, Q int
	fmt.Scanln(&N, &Q)

	r := ring.New(N)

	var L, R *ring.Ring
	r.Value = 1
	L = r
	for i := 2; i <= N; i++ {
		r = r.Next()
		r.Value = i

		if i == 2 {
			R = r
		}
	}

	var count int
	for i := 0; i < Q; i++ {
		var Opt string
		var No int
		fmt.Scanln(&Opt, &No)

		if Opt == "L" {
			l, c := Find(L, R, No)
			L = l
			count += c
		}

		if Opt == "R" {
			r, c := Find(R, L, No)
			R = r
			count += c
		}
	}

	fmt.Println(count)
}

func Find(master, sub *ring.Ring, n int) (*ring.Ring, int) {
	count := 0

	if master.Value == n {
		return master, count
	}

	for p := master.Next(); p != master; p = p.Next() {
		if p.Value == sub.Value {
			break
		}

		count++

		if p.Value == n {
			return p, count
		}
	}

	count = 0

	for p := master.Prev(); p != master; p = p.Prev() {
		if p.Value == sub.Value {
			break
		}

		count++
		if p.Value == n {
			return p, count
		}
	}

	panic("not found")
}
