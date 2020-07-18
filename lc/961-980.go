package lc

import (
	"fmt"
	"sort"
)

func subarraysDivByK(A []int, K int) int {
	divide := map[int]int{}
	rst := 0
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		mod := sum % K
		if mod > 0 {
			rst += divide[-(K-mod)%K]
		} else if mod < 0 {
			rst += divide[(K+mod)%K]
		}
		rst += divide[mod]
		if mod == 0 {
			rst++
		}
		divide[mod]++
	}
	return rst
}

func sortedSquares(A []int) []int {
	if len(A) == 1 {
		A[0] *= A[0]
		return A
	}
	k := sort.Search(len(A)-1, func(i int) bool {
		return A[i] <= A[i+1] && A[i] >= 0
	})
	if k == len(A) {
		k = len(A) - 1
	}
	dst := make([]int, len(A))
	i := k - 1
	j := k
	fmt.Println(k)
	for p := 0; i >= 0 || j < len(A); p++ {
		if j == len(A) || i >= 0 && A[i]*A[i] <= A[j]*A[j] {
			dst[p] = A[i] * A[i]
			i--
		} else {
			dst[p] = A[j] * A[j]
			j++
		}
	}
	return dst
}

type nodepath struct {
	x int
	y int
	l int
}

func uniquePathsIII(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var dir = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	var start nodepath
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				sum++
			}
			if grid[i][j] == 1 {
				start = nodepath{
					x: i,
					y: j,
				}
			}
		}
	}

	var getIndex = func(nd nodepath) int {
		return nd.x*n + nd.y
	}
	// check valid start
	var visited = make([]int, m*n)
	var stack []nodepath
	var queue []nodepath
	queue = append(queue, start)
	rst := 0
TRAVEL:
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for len(stack) > 0 && stack[len(stack)-1].l >= top.l {
			visited[getIndex(stack[len(stack)-1])] = 0
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, top)
		visited[getIndex(top)] = 1

		var next []nodepath

		for _, nb := range dir {
			x1, y1 := top.x+nb[0], top.y+nb[1]
			if x1 >= 0 && x1 < m && y1 >= 0 && y1 < n {
				v := grid[x1][y1]
				if v != -1 && visited[x1*n+y1] == 0 {
					if grid[x1][y1] == 2 && len(stack) == sum+1 {
						rst++
						goto TRAVEL
					}

					if v == 0 {
						next = append(next, nodepath{
							x: x1, y: y1, l: top.l + 1,
						})
					}
				}
			}
		}

		if len(next) > 0 {
			tmp := queue
			queue = append(next, tmp...)
		}
	}
	return rst
}
