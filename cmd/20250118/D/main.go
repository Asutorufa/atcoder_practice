package main

import "fmt"

func main() {
	var r int
	fmt.Scanln(&r)

	// 判断是否在圆内
	in := func(a, b int) bool {
		return (2*a+1)*(2*a+1)+(2*b+1)*(2*b+1) <= 4*r*r
	}

	cnt := 0
	up := r - 1
	res := (r-1)*4 + 1 // 圆心+直径上的方块

	// 只计算第一象限的数量，然后乘以四
	// 在应用数学里，平面直角坐标系中，右上角的象限称为第一象限。
	//
	// 圆心为(0, 0), 然后从x==1开始，直到大概x==r的位置
	for x := 1; in(x, 1); x++ {
		for !in(x, up) {
			up-- //找出在圆内的最高的那个方块，当不在圆内时，不断减小最上面的点的位置
		}

		cnt += up // 加上当前x的在圆内的方块
	}

	res += cnt * 4 // 四个象限，乘以四

	fmt.Println(res)
}
