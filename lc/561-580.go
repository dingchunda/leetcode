package lc

func findPaths(m int, n int, N int, i int, j int) int {
	memory := map[[3]int]int{}
	var dfs func(x, y, path int) int
	mod := int64(1000000007)
	dfs = func(x, y, path int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 1
		}
		if path == N {
			return 0
		}
		if v, ok := memory[[3]int{x, y, N - path}]; ok {
			return v
		}
		sum := int64(0)
		for _, v := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			sum = (sum + int64(dfs(x+v[0], y+v[1], path+1))) % mod
		}
		memory[[3]int{x, y, N - path}] = int(sum % mod)
		return int(sum % mod)
	}
	return dfs(i, j, 0)
}
