package lc

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := map[*TreeNode]int{}
	maxLen := 0
	var calHeight func(n *TreeNode) int
	calHeight = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		if v, ok := height[n]; ok {
			return v
		}
		h := max(calHeight(n.Left), calHeight(n.Right)) + 1
		height[n] = h
		l := calHeight(n.Left) + calHeight(n.Right)
		if l > maxLen {
			maxLen = l
		}
		// fmt.Println("node",n.Val,h,l)
		return h
	}
	calHeight(root)
	return maxLen
}

func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var dfs func(l, r, k int) int
	dfs = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		for r > l && boxes[r] == boxes[r-1] {
			r--
			k++
		}
		dp[l][r][k] = dfs(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = max(dp[l][r][k], dfs(l, i, k+1)+dfs(i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return dfs(0, len(boxes)-1, 0)
}

func subarraySum(nums []int, k int) int {
	rst := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for t := i; t >= 0; t-- {
			sum += nums[t]
			if sum == k {
				rst++
			}
		}
	}
	return rst
}

func findCircleNum(M [][]int) int {
	visitNodes := make([]bool, len(M))
	index := -1
	lowLink := make([]int, len(M))
	indexes := make([]int, len(M))
	for i := range indexes {
		indexes[i] = -1
	}
	ans := 0
	var dfs func(at int)
	dfs = func(at int) {
		index++
		visitNodes[at] = true
		border := M[at]
		lowLink[at] = index
		indexes[at] = index
		for j := 0; j < len(border); j++ {
			if at == j || border[j] == 0 {
				continue
			}
			if visitNodes[j] {
				lowLink[at] = min(lowLink[at], indexes[j])
			} else {
				dfs(j)
				lowLink[at] = min(lowLink[at], lowLink[j])
			}
		}
		if lowLink[at] == indexes[at] {
			ans++
		}
	}
	for i := 0; i < len(M); i++ {
		if indexes[i] == -1 {
			dfs(i)
		}
	}
	return ans
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col := len(matrix), len(matrix[0])
	ans := make([][]int, row)
	for i := range ans {
		ans[i] = make([]int, col)
	}
	var buf, tmp []int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			buf = buf[:0]
			tmp = tmp[:0]
			buf = append(buf, i*col+j)
			visited := map[int]bool{}
			dis := 0
			for len(buf) > 0 {
				tmp = tmp[:0]
				for _, n := range buf {
					if visited[n] {
						continue
					}
					visited[n] = true
					x, y := n/col, n%col
					if matrix[x][y] == 0 {
						ans[i][j] = dis
						goto next
					}
					if x > 0 {
						tmp = append(tmp, n-col)
					}
					if x < row-1 {
						tmp = append(tmp, n+col)
					}
					if y > 0 {
						tmp = append(tmp, n-1)
					}
					if y < col-1 {
						tmp = append(tmp, n+1)
					}
				}
				dis++
				buf, tmp = tmp, buf
			}
		next:
		}
	}
	return ans
}

func maxDepthMulti(root *NrTreeNode) int {
	var dfs func(r *NrTreeNode) int
	dfs = func(r *NrTreeNode) int {
		if r == nil {
			return 0
		}
		v := 0
		for _, n := range r.Children {
			v = max(v, dfs(n))
		}
		return v + 1
	}
	return dfs(root)
}
