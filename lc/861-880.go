package lc

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil {
		return nil
	}
	var rst = map[int]bool{}
	var child func(r *TreeNode, black *TreeNode, dis int)
	var buf, tmp []*TreeNode
	child = func(r *TreeNode, black *TreeNode, dis int) {
		buf = buf[:0]
		tmp = tmp[:0]
		buf = append(buf, r)
		lvl := 0
		for len(buf) > 0 && lvl < dis {
			lvl++
			for _, n := range buf {
				if n.Left != black && n.Left != nil {
					tmp = append(tmp, n.Left)
				}
				if n.Right != black && n.Right != nil {
					tmp = append(tmp, n.Right)
				}
			}
			buf, tmp = tmp, buf
			tmp = tmp[:0]
		}
		for _, n := range buf {
			rst[n.Val] = true
		}
	}
	find := false
	dis := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r == target {
			find = true
			child(target, nil, K)
			return
		}
		if !find {
			dfs(r.Left)
			if find {
				dis++
				if K-dis > 0 {
					child(r, r.Left, K-dis)
				}
				if dis == K {
					rst[r.Val] = true
				}
			}
		}

		if !find {
			dfs(r.Right)
			if find {
				dis++
				if K-dis > 0 {
					child(r, r.Right, K-dis)
				}
				if dis == K {
					rst[r.Val] = true
				}
			}
		}
	}
	dfs(root)
	var dup []int
	for k := range rst {
		dup = append(dup, k)
	}
	return dup
}

func lenLongestFibSubseq2(A []int) int {
	index := map[int]int{}
	for i, v := range A {
		index[v] = i
	}
	ans := 0
	nodes := map[int]int{}
	for k := 0; k < len(A); k++ {
		for j := k - 1; j >= 0; j-- {
			s := A[k] - A[j]
			i, ok := index[s]
			if ok {
				v, ok := nodes[i*len(A)+j]
				if !ok {
					v = 2
				}
				nodes[j*len(A)+k] = v + 1
				if v+1 > ans {
					ans = v + 1
				}
			}
		}
	}
	return ans
}

func lenLongestFibSubseq(A []int) int {
	index := map[int]int{}
	for i, v := range A {
		index[v] = i
	}
	ans := 0
	for s := 0; s < len(A); s++ {
		for j := s + 1; j < len(A); j++ {
			a0, a1 := A[s], A[j]
			sum := 2
			for {
				a3At, ok := index[a0+a1]
				if !ok {
					break
				}
				a0, a1 = a1, A[a3At]
				sum++
			}
			if sum > 2 && sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

func transpose(A [][]int) [][]int {
	if len(A) == 0 || len(A[0]) == 0 {
		return nil
	}
	ans := make([][]int, len(A[0]))
	for i := range ans {
		ans[i] = make([]int, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			ans[j][i] = A[i][j]
		}
	}
	return ans
}
