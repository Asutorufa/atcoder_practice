package main

import (
	"fmt"
	"strings"
)

/*
#include <bits/stdc++.h>
using namespace std;

	int main() {
		int n, m;
		cin >> n >> m;
		vector<vector<int>> w = { {} };
		for (int i = 1; i <= n; i++) {
			vector<vector<int>> v;
			for (vector<int> a : w) {
				for (int x = (i == 1 ? 1 : a.back() + 10); x <= m - 10 * (n - i); x++) {
					vector<int> na = a;
					na.push_back(x);
					v.push_back(na);
				}
			}
			swap(v, w);
		}
		int X = w.size();
		cout << X << '\n';
		for (int i = 0; i < X; i++) for (int j = 0; j < n; j++) cout << w[i][j] << " \n"[j == n - 1];
	}
*/

/*

 */

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	w := [][]int{{}}

	// 由计算过程可知，每个循环添加当前位置的元素直接到当前元素所能达到的最大值
	// [[1] [2] [3]]
	// [[1 11] [1 12] [1 13] [2 12] [2 13] [3 13]]
	// [[1 11 21] [1 11 22] [1 11 23] [1 12 22] [1 12 23] [1 13 23] [2 12 22] [2 12 23] [2 13 23] [3 13 23]]
	for i := 1; i <= N; i++ {
		v := [][]int{}
		for _, a := range w {
			// M-(10*(N-i)) 为当前元素的最大值
			//  如例子中的 3 23:
			//   第一位元素，则 M-(10*(N-i)) = 23-(10*(3-1)) = 3, 第一位最大为3
			//   第二位元素，则 M-(10*(N-i)) = 23-(10*(3-2)) = 13，第二位最大为13
			//   第三位元素，则 M-(10*(N-i)) = 23-(10*(3-3)) = 23，第三位最大为23
			// 由此类推，当N为其他位数时也可行
			//
			// 而当前位置的最小元素可由 or(i == 1, 1, back(a)+10) 所得
			//  如例子中的 3 23:
			//   第一位元素，因为之前没有添加过任何元素, 第一位最小为1，最后结果为 [[1] [2] [3]]
			//   第二位元素
			//    当第一位为1时，最小为11，可得 [[1 11] [1 12] [1 13] [2] [3]]
			//    当第一位为2时，最小为12，可得 [[1 11] [1 12] [1 13] [2 12] [2 13] [3]]
			//    当第一位为3时，最小为13，可得 [[1 11] [1 12] [1 13] [2 12] [2 13] [3 13]]
			//  由此类推可得到最终结果
			for x := or(i == 1, 1, back(a)+10); x <= M-(10*(N-i)); x++ {
				// 如果不复制，会导致扩容之后，把之前的值覆盖掉
				na := append([]int(nil), a...) // 配列をコピー
				v = append(v, append(na, x))

				v = append(v, append(a, x))
			}
		}

		w = v
	}

	fmt.Println(len(w))

	/*
		10
		1 11 21
		1 11 22
		1 11 23
		1 12 22
		1 12 23
		1 13 23
		2 12 22
		2 12 23
		2 13 23
		3 13 23
	*/
	for _, v := range w {
		fmt.Println(strings.Trim(fmt.Sprint(v), "[]"))
	}
}

func back(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[len(a)-1]
}

func or[T any](x bool, a, b T) T {
	if x {
		return a
	}
	return b
}
