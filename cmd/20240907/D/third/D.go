package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
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

	remain := h * w

	harray := make([]*set.Set[int], h)
	warray := make([]*set.Set[int], w)

	for i := range harray {
		for j := range warray {
			if harray[i] == nil {
				harray[i] = set.New(comparator.IntComparator, set.WithGoroutineSafe())
				harray[i].Insert(-1)
				harray[i].Insert(w)
			}
			if warray[j] == nil {
				warray[j] = set.New(comparator.IntComparator, set.WithGoroutineSafe())
				warray[j].Insert(-1)
				warray[j].Insert(h)
			}

			harray[i].Insert(j)
			warray[j].Insert(i)
		}
	}

	erase := func(i, j int) {
		harray[i].Erase(j)
		warray[j].Erase(i)
	}

	for qi := 0; qi < q; qi++ {
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

		{
			it := harray[x].LowerBound(y)
			if it.Value() == y {
				erase(x, y)
				remain--
				continue
			} else {
				if it.Value() != w {
					erase(x, it.Value())
					remain--
				}
				it = harray[x].LowerBound(y)
				it.Prev()
				if it.Value() != -1 {
					erase(x, it.Value())
					remain--
				}
			}
		}
		{
			it := warray[y].LowerBound(x)
			if it.Value() != h {
				erase(it.Value(), y)
				remain--
			}
			it = warray[y].LowerBound(x)
			it.Prev()
			if it.Value() != -1 {
				erase(it.Value(), y)
				remain--
			}
		}
	}

	fmt.Println(remain)
}
