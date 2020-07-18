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