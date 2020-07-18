package lc

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
