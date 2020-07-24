package lc

func networkDelayTime(times [][]int, N int, K int) int {
	shortest := map[int]int{}
	borders := make([][][2]int, N+1)
	for _, t := range times {
		borders[t[0]] = append(borders[t[0]], [2]int{t[1], t[2]})
	}
	visited := make([]bool, N+1)
	shortest[K] = 0
	hit := 0
	ans := 0
	for len(shortest) > 0 {
		dis := 100000
		top := -1
		for k, v := range shortest {
			if v < dis {
				dis = v
				top = k
			}
		}

		delete(shortest, top)
		visited[top] = true
		hit++
		ans = dis
		for _, n := range borders[top] {
			if visited[n[0]] {
				continue
			}
			nextDis := dis + n[1]
			weight, ok := shortest[n[0]]
			if !ok || nextDis < weight {
				shortest[n[0]] = nextDis
			}
		}
	}
	if hit == N {
		return ans
	}
	return -1
}
