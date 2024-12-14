package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	var N, S int
	fmt.Fscanln(br, &N, &S)

	var As []int
	var sum int
	readIntFunc(N, br, func(x int, index int) {
		As = append(As, x)
		sum += x
	})

	/*
		想复杂了，还上了树状数组
		其实只需要动态规划的思想，这方面太弱鸡了

			3 42
			3 8 4

			checker = S - (S % sum) = 42 - (42 % 15) = 42 - 12 = 30

			30 + 3 = 33
			33 + 8 = 41
			41 + 4 = 45

			45 - 3 = 42

			一直计算直到大于S, 然后判断是否等于S
			如果大于S就把整个序列往后移一位，如上面的例子往后移一位减去3

			然后循环上面的操作
	*/
	checker := S - (S % sum)

	var j int
	for _, v := range As {
		for checker < S {
			checker += As[j%N]
			j++
		}

		if checker == S {
			fmt.Println("Yes")
			return
		}

		checker -= v
	}

	fmt.Println("No")
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}

type Result struct {
	Count  int
	Indexs []int
	Sum    int
}

func Sums2(a []int) map[int]Result {
	tree := make([]int, len(a))
	copy(tree, a)

	initBIT(tree)

	res := map[int]Result{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j+i >= len(a) {
				break
			}

			sum := intervalSum(tree, j, j+i)

			res[sum] = Result{Count: i, Indexs: []int{j, j + i}, Sum: sum}
		}
	}
	return res
}

func Sums3(a []int) map[int]Result {
	tree := make([]int, len(a))
	copy(tree, a)

	initBIT(tree)

	res := map[int]Result{}
	for i := range a {
		sum := intervalSum(tree, 0, i)

		res[sum] = Result{Count: i, Indexs: []int{0, i}, Sum: sum}
	}

	for i := range a {
		sum := intervalSum(tree, len(a)-1-i, len(a)-1)

		res[sum] = Result{Count: i, Indexs: []int{len(a) - 1 - i, len(a) - 1}, Sum: sum}
	}

	return res
}

func SumsFind(a []int, k int) bool {
	tree := make([]int, len(a))
	copy(tree, a)

	initBIT(tree)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j+i >= len(a) {
				break
			}

			sum := intervalSum(tree, j, j+i)

			if sum == k {
				return true
			}

		}
	}

	return false
}

func SumsFind2(a []int, k int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j+i >= len(a) {
				break
			}

			sum := intervalSum(a, j, j+i)

			if sum == k {
				return true
			}

		}
	}

	return false
}

func Sums(a []int) map[int]Result {
	res := map[int]Result{}
	for i := 1; i <= len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j+i > len(a) {
				break
			}

			sum := Sum(a[j : j+i])
			res[sum] = Result{Count: i, Indexs: []int{j, j + i}, Sum: sum}
		}
	}
	return res
}

func Sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func lowbit(x int) int {
	return x & -x
}

// 将arr[idx]的值加上val
func add(arr []int, idx, val int) {
	idx++
	for idx <= len(arr) {
		arr[idx-1] += val
		idx += lowbit(idx)
	}
}

// 求[0,idx]的前缀和
func prefixSum(arr []int, idx int) int {
	var ans int
	idx++
	for idx != 0 {
		ans += arr[idx-1]
		idx -= lowbit(idx)
	}
	return ans
}

// 求[i,j]的区间和
func intervalSum(arr []int, i, j int) int {
	return prefixSum(arr, j) - prefixSum(arr, i-1)
}

// 在原数组上初始化树状数组
func initBIT(arr []int) {
	for i := len(arr) - 2; i >= 0; i-- {
		val := arr[i]
		arr[i] = 0
		add(arr, i, val)
	}
}

// func main() {
// 	arr := []int{4, 2, 7, 5, 9, 1, 0, 3}
// 	initBIT(arr)
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(prefixSum(arr, i))
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			fmt.Println(i, j, intervalSum(arr, i, j))
// 		}
// 	}
// }
