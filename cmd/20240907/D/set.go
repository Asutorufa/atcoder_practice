package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
)

type TreaSet struct {
	root *TreaSetNode
}

func (t *TreaSet) Insert(k int) {
	if t == nil {
		t = &TreaSet{}
	}
	t.root = t.root.insert(k, rand.Int())
}

func (t *TreaSet) Find(k int) bool {
	if t == nil {
		return false
	}
	return t.root.find(k)
}

func (t *TreaSet) Delete(k int) {
	if t == nil {
		return
	}
	t.root = t.root.delete(k)
}

func (t *TreaSet) InOrder(wr io.Writer) {
	if t == nil {
		return
	}
	t.root.inOrder(wr)
}

func (t *TreaSet) LowerBound(k int) (int, bool) {
	if t == nil {
		return k, false
	}
	return t.root.lowerBound(k)
}

func (t *TreaSet) InverseLowerBound(k int) (int, bool) {
	if t == nil {
		return k, false
	}
	return t.root.inverseLowerBound(k)
}

func (t *TreaSet) Size() int {
	if t == nil {
		return 0
	}
	return t.root.size()
}

type TreaSetNode struct {
	key         int
	priority    int
	left, right *TreaSetNode
}

func (nd *TreaSetNode) insert(k int, p int) *TreaSetNode {
	if nd == nil {
		return &TreaSetNode{key: k, priority: p, left: nil, right: nil}
	}
	if k == nd.key {
		return nd
	} else if nd.key > k {
		nd.left = nd.left.insert(k, p)
		if nd.priority < nd.left.priority {
			nd = nd.rightRotate()
		}
	} else {
		nd.right = nd.right.insert(k, p)
		if nd.priority < nd.right.priority {
			nd = nd.leftRotate()
		}
	}
	return nd
}

func (nd *TreaSetNode) find(k int) bool {
	result := false
	if nd == nil {
		return result
	}
	if nd.key == k {
		result = true
	} else if nd.key > k {
		result = nd.left.find(k)
	} else {
		result = nd.right.find(k)
	}
	return result
}

func (nd *TreaSetNode) delete(k int) *TreaSetNode {
	if nd == nil {
		return nil
	}
	if nd.key > k {
		nd.left = nd.left.delete(k)
	} else if nd.key < k {
		nd.right = nd.right.delete(k)
	} else {
		if nd.left == nil && nd.right == nil {
			return nil
		} else if nd.left == nil {
			nd = nd.leftRotate()
		} else if nd.right == nil {
			nd = nd.rightRotate()
		} else {
			if nd.left.priority > nd.right.priority {
				nd = nd.rightRotate()
			} else {
				nd = nd.leftRotate()
			}
		}
		return nd.delete(k)
	}
	return nd
}

func (nd *TreaSetNode) rightRotate() *TreaSetNode {
	s := nd.left
	nd.left = s.right
	s.right = nd
	return s
}

func (nd *TreaSetNode) leftRotate() *TreaSetNode {
	s := nd.right
	nd.right = s.left
	s.left = nd
	return s
}

func (nd *TreaSetNode) inOrder(wr io.Writer) {
	if nd != nil {
		nd.left.inOrder(wr)
		fmt.Fprintf(wr, " %v", nd.key)
		nd.right.inOrder(wr)
	}
}

func (nd *TreaSetNode) lowerBound(k int) (int, bool) {
	if nd == nil {
		return k, false
	}
	result := make([]int, 0)
	if nd.key >= k {
		result = append(result, nd.key)
		result1, ok1 := nd.left.lowerBound(k)
		if ok1 {
			result = append(result, result1)
		}
	} else {
		result2, ok2 := nd.right.lowerBound(k)
		if ok2 {
			result = append(result, result2)
		}
	}
	if len(result) == 0 {
		return nd.key, false
	} else {
		return Min(result...), true
	}
}

func (nd *TreaSetNode) inverseLowerBound(k int) (int, bool) {
	if nd == nil {
		return k, false
	}
	result := make([]int, 0)
	if nd.key <= k {
		result = append(result, nd.key)
		result2, ok2 := nd.right.inverseLowerBound(k)
		if ok2 {
			result = append(result, result2)
		}
	} else {
		result1, ok1 := nd.left.inverseLowerBound(k)
		if ok1 {
			result = append(result, result1)
		}
	}
	if len(result) == 0 {
		return nd.key, false
	} else {
		return Max(result...), true
	}
}

func (nd *TreaSetNode) size() int {
	if nd == nil {
		return 0
	}
	return nd.left.size() + nd.right.size() + 1
}

func Max(num ...int) int {
	max := num[0]
	for i := 1; i < len(num); i++ {
		if max < num[i] {
			max = num[i]
		}
	}
	return max
}

func Min(num ...int) int {
	min := num[0]
	for i := 1; i < len(num); i++ {
		if min > num[i] {
			min = num[i]
		}
	}
	return min
}

var inf = 1 << 60

// st := TreeSet(n+10)
type st_type struct {
	sz     int
	maxval int
	dat    []int
	mp     map[int]int
}

func TreeSet(n int) *st_type {
	sz, maxval := 0, n+1
	dat := make([]int, maxval) // 1-index : 1 ~ n
	mp := make(map[int]int)
	return &st_type{sz: sz, maxval: maxval, dat: dat, mp: mp}
}

func (st *st_type) Add(val int) {
	for i := val; i < st.maxval; i += (i & -i) {
		st.dat[i]++
	}
	st.sz++
	st.mp[val]++
}

func (st *st_type) Remove(val int) {
	if !st.Find(val) {
		return
	}
	for i := val; i < st.maxval; i += (i & -i) {
		st.dat[i]--
	}
	st.sz--
	st.mp[val]--
}

func (st *st_type) Find(val int) bool {
	s1, s2 := 0, 0
	for i := val; i > 0; i -= (i & -i) {
		s2 += st.dat[i]
	}
	for i := val - 1; i > 0; i -= (i & -i) {
		s1 += st.dat[i]
	}
	if s2-s1 > 0 {
		return true
	} else {
		return false
	}
}

func (st *st_type) Size() int {
	return st.sz
}

func (st *st_type) Show() { // for debug
	vals := make([]int, 0)
	for key, val := range st.mp {
		for i := 0; i < val; i++ {
			vals = append(vals, key)
		}
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})
	fmt.Fprintln(os.Stdout, vals)
}

func (st *st_type) KthValue(k int) int {
	if k <= 0 { // 1-index
		return -inf
	}
	if k > st.sz {
		return inf
	}
	ret, start := 0, 1
	for {
		if (start << 1) > st.maxval {
			break
		}
		start <<= 1
	}
	for i := start; i > 0; i >>= 1 {
		if ret+i < st.maxval && st.dat[ret+i] < k {
			k -= st.dat[ret+i]
			ret += i
		}
	}
	return ret + 1 // 1-index
}

func (st *st_type) LowerBound(val int) int {
	if val > st.maxval-1 {
		return inf
	}
	s1 := 0
	for i := val - 1; i > 0; i -= (i & -i) {
		s1 += st.dat[i]
	}
	return st.KthValue(s1 + 1)
}

func (st *st_type) ReverseLowerBound(val int) int {
	if val > st.maxval-1 {
		val = st.maxval - 1
	}
	s1 := 0
	for i := val; i > 0; i -= (i & -i) {
		s1 += st.dat[i]
	}
	if s1 == 0 {
		return -inf
	} else {
		return st.KthValue(s1)
	}
}
