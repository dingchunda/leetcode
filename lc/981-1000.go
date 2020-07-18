package lc

func findJudge(N int, trust [][]int) int {
	table := make([][]bool, N)
	for i := range table {
		table[i] = make([]bool, N)
	}
	for _, t := range trust {
		table[t[0]-1][t[1]-1] = true
	}
	//fmt.Println(table)
	for i := 0; i < len(table); i++ {
		ok := true
		for j := 0; j < len(table); j++ {
			if i != j {
				if !table[j][i] {
					ok = false
					break
				}
			}
		}
		if ok {
			for k := 0; k < N; k++ {
				if k != i {
					if table[i][k] {
						ok = false
						break
					}
				}
			}
			if ok {
				return i + 1
			}
		}
	}
	return -1
}
