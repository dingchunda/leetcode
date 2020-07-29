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

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var dfs func(r *TreeNode)
	var ans string
	var path []byte
	dfs = func(r *TreeNode) {
		if r.Left == nil && r.Right == nil {
			dst := make([]byte, len(path))
			copy(dst, path)
			i, j := 0, len(path)-1
			for i < j {
				dst[i], dst[j] = dst[j], dst[i]
				i++
				j--
			}
			if s := string(dst); ans == "" || s < ans {
				ans = s
			}
		}
		for _, n := range []*TreeNode{r.Left, r.Right} {
			if n != nil {
				path = append(path, byte(n.Val+'a'))
				dfs(n)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(root)
	return ""
}
