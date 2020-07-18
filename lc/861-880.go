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
