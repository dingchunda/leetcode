package lc

import "math"

func isBipartite(graph [][]int) bool {
	visited := map[int]int{}
	var dfs func(at int) bool
	dfs = func(at int) bool {
		c := visited[at]
		for _, next := range graph[at] {
			color, ok := visited[next]
			if ok && color == c {
				return false
			}
			if !ok {
				visited[next] = 3 - c%3
				if !dfs(next) {
					return false
				}
			}
		}
		return true
	}

	for i := range graph {
		if _, ok := visited[i]; ok {
			continue
		}
		visited[i] = 1
		if !dfs(i) {
			return false
		}
	}
	return true
}

func minDiffInBST(root *TreeNode) int {
	var travel func(n *TreeNode, visitor func(v int))
	travel = func(n *TreeNode, visitor func(v int)) {
		if n == nil {
			return
		}
		travel(n.Left, visitor)
		visitor(n.Val)
		travel(n.Right, visitor)
	}
	last := -1
	ans := math.MaxInt32
	travel(root, func(v int) {
		if last > 0 {
			if diff := v - last; diff < ans {
				ans = diff
			}
		}
		last = v
	})
	return ans
}
