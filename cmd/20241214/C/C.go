package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a, b, c, d, e int
	fmt.Scanln(&a, &b, &c, &d, &e)

	maps := map[rune]int{
		'A': a,
		'B': b,
		'C': c,
		'D': d,
		'E': e,
	}

	var res results
	for _, s := range ALL {
		score := 0

		for _, v := range s {
			score += maps[v]
		}

		res = append(res, Result{S: s, Score: score})
	}

	sort.Sort(res)

	for _, r := range res {
		fmt.Println(r.S)
	}
}

type results []Result

func (r results) Len() int {
	return len(r)
}
func (r results) Less(i, j int) bool {
	if r[i].Score == r[j].Score {
		return r[i].S < r[j].S
	}

	return r[i].Score > r[j].Score
}
func (r results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type Result struct {
	S     string
	Score int
}

var ALL = strings.Split(`ABCDE
ACDE
BCDE
ABCE
ABDE
ABCD
CDE
ACE
ADE
BCE
BDE
ABE
ACD
BCD
ABC
ABD
CE
DE
AE
BE
CD
AC
AD
BC
BD
AB
E
C
D
A
B`, "\n")
