package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
)

/*
依旧是贪心算法？

#include <bits/stdc++.h>
using namespace std;

	int main() {
	    int n;
	    string s;
	    cin >> n >> s;
	    int c1 = 0;
	    for (char c: s) {
	        if (c == '1') ++c1;
	    }
	    int now = 0;
	    long long ans = 0;
	    for (char c: s) {
	        if (c == '0') {
	            ans += min(now, c1 - now);
	        } else {
	            ++now;
	        }
	    }
	    cout << ans << endl;
	}

c1: 1的总数
now: 当前遍历过1的个数
c1 - now: 1的总数减去当前遍历过1的数量
min(now, c1 - now): 取二者中的较小值，表示这个 '0' 能够配对的 '1' 的数量（类似经过这个0的1的数量？）。

每个0能够配对的1的数量加起来就是所有1需要移动的步数。
*/
func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	var S string
	fmt.Fscanln(br, &S)

	start := strings.IndexByte(S, '1')
	end := strings.LastIndexByte(S, '1')

	S = S[start : end+1]

	if len(S) <= 2 {
		fmt.Println(0)
		return
	}

	// 这个是直接把中间的1作为目标，感觉可能不太对
	ones := oneIndex(S)
	mid := len(ones) / 2
	target := ones[mid]

	var moves int

	for i := 0; i < mid; i++ {
		// 计算当前节点移动到目标节点需要多少步
		// target - i   中间位置（目标）减去当前位置
		//
		// 然后再加上 这个节点本身的位置 mid-i
		// 这里mid - i 是因为当前节点到目标中间的1因为已经被移到前面了，所以需要减去这一部分的

		moves += target - mid + i - ones[i]
	}

	// 后半部分刚好反过来
	for i := mid + 1; i < len(ones); i++ {
		moves += ones[i] - (target + i - mid)
	}

	fmt.Println(moves)
}

func oneIndex(s string) []int {
	var res []int
	for i, v := range s {
		if v == '1' {
			res = append(res, i)
		}
	}

	return res
}

func main2() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	var S string
	fmt.Fscanln(br, &S)

	start := strings.IndexByte(S, '1')
	end := strings.LastIndexByte(S, '1')

	S = S[start : end+1]

	if len(S) <= 2 {
		fmt.Println(0)
		return
	}

	oneIndexs := oneIndex(S)

	var current int = -1
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, v := range oneIndexs {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			// fmt.Println(S[:v], S[v:])

			front := []byte(S[:v])
			Reverse(front)
			back := []byte(S[v:])

			wg := sync.WaitGroup{}
			wg.Add(2)

			var fc, bc int
			go func() {
				defer wg.Done()
				bc = count(back)
			}()
			go func() {
				defer wg.Done()
				fc = count(front)
			}()

			wg.Wait()

			mu.Lock()
			defer mu.Unlock()

			if current == -1 {
				current = bc + fc
				return
			}

			if current > bc+fc {
				current = bc + fc
			}

		}(v)
	}

	wg.Wait()
	// fmt.Println(S[:mid], S[mid:])
	// fmt.Println(S, start, end)

	// fmt.Println(bc+fc, bc, fc)

	fmt.Println(current)
}

// Reverse reverses the elements of the slice in place.
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func count(back []byte) int {
	var emptyCount int
	var bc int
	for len(back) > 0 {
		if back[0] == '1' {
			back = back[1:]

			if emptyCount > 0 {
				bc += emptyCount
			}

			continue
		}

		i := bytes.IndexByte(back, '1')
		if i == -1 {
			break
		}

		bc += i + emptyCount
		emptyCount += i
		// fmt.Println(i, emptyCount, bc, string(back))

		back = back[i+1:]

	}

	return bc
}

func countN(back []byte) int {
	var bc int
	for len(back) > 0 {

		if back[0] == '1' {
			back = back[1:]
			continue
		}

		i := bytes.IndexByte(back, '1')
		if i == -1 {
			break
		}

		bc += i
		back[i] = '0'
		back = back[1:]

	}

	return bc
}
