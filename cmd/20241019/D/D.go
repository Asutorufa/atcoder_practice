package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscanln(r, &N, &M)

	graph := make([][]int, N)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)

		graph[a-1] = append(graph[a-1], b-1)
	}

	visited := make([]int, N)
	for i := 0; i < N; i++ {
		visited[i] = -1
	}

	visited[0] = 0
	queue := []int{0}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, c := range graph[v] {
			if c == 0 {
				visited[c] = visited[v] + 1
				fmt.Println(visited[0])
				return
			} else {
				if visited[c] <= 0 {
					visited[c] = visited[v] + 1
					queue = append(queue, c)
				}
			}
		}
	}

	fmt.Println(-1)
}

/*
6 9
6 1
1 5
2 6
2 1
3 6
4 2
6 4
3 5
5 4

graph:　[[4] [5 0] [5 4] [1] [3] [0 3]]

queue: 0
v: 0, p: [4]
c:4 v[4] = v[0] + 1 = 0 + 1 = 1

queue: 4
v: 4, p: [3]
c:3 v[3] = v[4] + 1 = 1 + 1 = 2

queue: 3
v: 3, p: [1]
c:1 v[1] = v[3] + 1 = 2 + 1 = 3

queue: 1
v: 1, p: [5, 0]
c:5 v[5] = v[1] + 1 = 3 + 1 = 4

c:0 v[0] = v[1] + 1 = 3 + 1 = 4

c == 0, so: result = 4
*/
